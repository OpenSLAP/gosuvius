package parser

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Much of this is blatantly stolen from https://github.com/golang/go/blob/07b8011393dc3d3a78b3cd0857a31da339985994/src/text/template/parse/lex.go
// Thanks, Rob!

type stateFn func(*lexer) stateFn

type itemType int
type item struct {
	typ itemType // The type of this item.
	pos int      // The starting position, in bytes, of this item in the input string.
	val string   // The value of this item.
}

func (i item) String() string {
	switch {
	case i.typ == itemEOF:
		return "EOF"
	case i.typ == itemError:
		return i.val
	}
	return fmt.Sprintf("%q", i.val)
}

type lexer struct {
	name    string    // the name of the input; used only for error reports
	input   string    // the string being scanned
	state   stateFn   // the next lexing function to enter
	pos     int       // current position in the input
	start   int       // start position of this item
	width   int       // width of last rune read from input
	lastPos int       // position of most recent item returned by nextItem
	items   chan item // channel of scanned items
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.start, l.input[l.start:l.pos]}
	l.start = l.pos
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, l.next()) {
	}
	l.backup()
}

// return true & increase l.pos if found, return false if not
func (l *lexer) acceptSequence(valid string) bool {
	if strings.HasPrefix(l.input[l.pos:], valid) {
		l.pos += len(valid)
		return true
	}
	return false
}

// accept until we hit 'end' or eof
// return true if we accept something, false otherwise
func (l *lexer) acceptUntil(end rune) bool {
	if l.peek() == end || l.peek() == eof {
		return false
	}
	for r := l.next(); r != end && r != eof; r = l.next() {
	}
	l.backup()
	return true
}

// lineNumber reports which line we're on, based on the position of
// the previous item returned by nextItem. Doing it this way
// means we don't have to worry about peek double counting.
func (l *lexer) lineNumber() int {
	return 1 + strings.Count(l.input[:l.lastPos], "\n")
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{itemError, l.start, fmt.Sprintf(format, args...)}
	return nil
}

// nextItem returns the next item from the input.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) nextItem() item {
	item := <-l.items
	l.lastPos = item.pos
	return item
}

// drain drains the output so the lexing goroutine will exit.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) drain() {
	for range l.items {
	}
}

// run runs the state machine for the lexer.
func (l *lexer) run() {
	for l.state = lexLine; l.state != nil; {
		l.state = l.state(l)
	}
	close(l.items)
}

// lex a number
func lexNumber(l *lexer) stateFn {
	if l.accept("-") {
		// we have either a "-" or numeric characters
		if !unicode.IsNumber(l.peek()) {
			l.emit(itemError)
			return nil
		}
	}
	l.acceptRun(digits)
	l.accept(".")
	l.acceptRun(digits)
	return lexLine
}

func lexCode(l *lexer) stateFn {
	r := l.next()
	l.accept(digits)
	l.emit(LeaderMapping[r])
	return lexSpace
}

func lexParam(l *lexer) stateFn {
	return nil
}

func lexComment(l *lexer) stateFn {
	l.acceptUntil('\n')
	l.emit(itemComment)
	return lexLine
}

// lexSpace scans a run of space characters.
// One space has already been seen.
func lexSpace(l *lexer) stateFn {
	for isSpace(l.peek()) {
		l.next()
	}
	l.emit(itemSpace)
	return lexLine
}

// assume default = garbage we don't care about
func lexLine(l *lexer) stateFn {
	r := l.next()
	if _, ok := SupportedCodes[r]; ok {
		l.backup()
		return lexCode
	} else if _, ok := SupportedParameters[r]; ok {
		l.backup()
		return lexParam
	} else if r == ';' {
		return lexComment
	} else if isSpace(r) {
		return lexSpace
	} else if r == '\n' {
		l.emit(itemNewline)
		return lexLine
	} else if r == eof {
		l.emit(itemEOF)
		return nil
	}

	l.errorf("FUCK ME! WHAT IS '%c'?\n", r)
	return nil
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

func isASCIIAlpha(r rune) bool {
	return unicode.IsLetter(r) && r < unicode.MaxASCII
}

// isAlphaNumeric reports whether r is in ALPHANUM_CHARS (English letters, Arabic numbers)
func isAlphaNumeric(r rune) bool {
	return isASCIIAlpha(r) || unicode.IsNumber(r)
}

// lex creates a new scanner for the input string.
func lex(name, input string) *lexer {
	l := &lexer{
		name:  name,
		input: input,
		items: make(chan item),
	}
	go l.run()
	return l
}
