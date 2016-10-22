package main

import (
	"bufio"
	"bytes"
	"io"
)

type Token int

const (
	// Misc
	ILLEGAL Token = iota
	EOF
	WHITESPACE
	NEWLINE
	COMMENT

	// Codes
	GCODE
	MCODE

	// Params
	EPARAM
	FPARAM
	SPARAM
	XPARAM
	YPARAM
	ZPARAM
)

var types_map = map[rune]Token{
	'g': GCODE,
	'G': GCODE,
	'm': MCODE,
	'M': MCODE,
	'e': EPARAM,
	'E': EPARAM,
	'f': FPARAM,
	'F': FPARAM,
	's': SPARAM,
	'S': SPARAM,
	'x': XPARAM,
	'X': XPARAM,
	'y': YPARAM,
	'Y': YPARAM,
	'z': ZPARAM,
	'Z': ZPARAM,
}

var types_reverse_map = map[Token]rune{
	GCODE:  'G',
	MCODE:  'M',
	EPARAM: 'E',
	FPARAM: 'F',
	SPARAM: 'S',
	XPARAM: 'X',
	YPARAM: 'Y',
	ZPARAM: 'Z',
}

var eof = rune(0)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	char, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return char
}

func (s *Scanner) unread() {
	err := s.r.UnreadRune()
	if err != nil {
		panic("derp")
	}
}

func (s *Scanner) scan() (tok Token, raw string) {
	char := s.read()

	if is_white_space(char) {
		if char == '\n' {
			return NEWLINE, string(char)
		} else {
			s.unread()
			return s.scan_whitespace()
		}
	} else if is_leader(char) {
		s.unread()
		return s.scan_ident()
	} else if char == ';' {
		s.unread()
		return s.scan_comment()
	}

	switch char {
	case eof:
		return EOF, ""
	}

	return ILLEGAL, string(char)
}

func (s *Scanner) scan_whitespace() (tok Token, raw string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if char := s.read(); char == eof {
			break
		} else if !is_white_space(char) {
			s.unread()
			break
		} else {
			buf.WriteRune(char)
		}
	}
	return WHITESPACE, buf.String()
}

func (s *Scanner) scan_ident() (tok Token, raw string) {
	var buf bytes.Buffer
	first := s.read()
	//buf.WriteRune(first)

	for {
		if char := s.read(); char == eof {
			break
		} else if !is_legit(char) {
			s.unread()
			break
		} else {
			buf.WriteRune(char)
		}
	}

	return types_map[first], buf.String()
}

func (s *Scanner) scan_comment() (tok Token, raw string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if char := s.read(); char == eof {
			break
		} else if char == '\n' {
			s.unread()
			break
		} else {
			buf.WriteRune(char)
		}
	}

	return COMMENT, buf.String()
}

func is_white_space(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n'
}

func is_op_leader(char rune) bool {
	return char == 'g' || char == 'G' || char == 'm' || char == 'M'
}

func is_xyz(char rune) bool {
	return (char >= 'x' && char <= 'z') || (char >= 'X' && char <= 'Z')
}

func is_param_leader(char rune) bool {
	return is_xyz(char) || char == 'e' || char == 'E' || char == 's' || char == 'S' || char == 'f' || char == 'F'
}

func is_leader(char rune) bool {
	if _, ok := types_map[char]; ok {
		return true
	}
	return false
}

func is_numeric(char rune) bool {
	return char >= '0' && char <= '9'
}

func is_letter(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func is_legit(char rune) bool {
	return is_letter(char) || is_numeric(char) || char == '.' || char == '-'
}

func is_code_token(t Token) bool {
	return t == GCODE || t == MCODE
}

func is_param_token(t Token) bool {
	return t > MCODE
}
