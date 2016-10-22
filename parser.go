package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type Line struct {
	CodeType   Token
	CodeNumber int
	RawCode    string
	Params     []Param
}

type Param struct {
	ParamType  Token
	RawParam   string
	FloatValue float64
}

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		raw string
		n   int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) scan() (tok Token, raw string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.raw
	}

	tok, raw = p.s.scan()
	p.buf.tok = tok
	p.buf.raw = raw

	return
}

func (p *Parser) unscan() {
	p.buf.n = 1
}

func (p *Parser) scan_ignore_unused() (tok Token, raw string) {
	// Collects tokens, skips comments and whitespace
	tok, raw = p.scan()
	if tok == WHITESPACE || tok == COMMENT || tok == NEWLINE {
		tok, raw = p.scan_ignore_unused()
	}

	return
}

func (p *Parser) Parse() (*Line, error) {
	l := &Line{}

	tok, raw := p.scan_ignore_unused()
	if tok == EOF {
		return nil, nil
	} else if tok != GCODE && tok != MCODE {
		fmt.Printf("Got an unexpected token... %d: %s", tok, raw)
		return nil, nil
	} else {
		int_val, _ := strconv.Atoi(raw)
		l.CodeNumber = int_val
		l.CodeType = tok
		l.RawCode = fmt.Sprintf("%c%s", types_reverse_map[tok], raw)
	}

	for {
		tok, raw := p.scan_ignore_unused()
		if is_code_token(tok) {
			p.unscan()
			break
		} else if is_param_token(tok) {
			if !is_valid_number(raw) {
				fmt.Printf("Invalid parameter value: %s\n", raw)
			}

			float_val, _ := strconv.ParseFloat(raw, 64)

			temp_param := Param{
				ParamType:  tok,
				RawParam:   fmt.Sprintf("%c%s", types_reverse_map[tok], raw),
				FloatValue: float_val,
			}
			l.Params = append(l.Params, temp_param)

		} else if tok == NEWLINE {
			break
		} else {
			p.unscan()
			break
		}
	}
	return l, nil
}

func is_valid_number(raw string) bool {
	period_found := false
	buf := bytes.NewBufferString(raw)
	first, _, _ := buf.ReadRune()

	if first != '-' {
		buf.UnreadRune()
	}

	for {
		if char, _, _ := buf.ReadRune(); char == eof {
			return true
		} else if !is_numeric(char) && char != '.' {
			return false
		} else if char == '.' {
			if period_found {
				return false
			} else {
				period_found = true
			}
		}
	}
	return true
}
