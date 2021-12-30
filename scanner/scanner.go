package scanner

import (
	"fmt"

	"github.com/weiser/lox/token"
)

//TODO: start pg 49, section 4.6

type Scanner struct {
	Source               string
	Tokens               []token.Token
	Errors               []Error
	Start, Current, Line int
}

type Error struct {
	Source               string
	Start, Current, Line int
}

func (e Error) String() string {
	return fmt.Sprintf("source=%v, Start=%v, current=%v, line=%v", e.Source, e.Start, e.Current, e.Line)
}

func MakeScanner(src string) Scanner {
	return Scanner{Source: src, Tokens: make([]token.Token, 0), Start: 0, Current: 0, Line: 1}
}

func (s *Scanner) ScanTokens() []token.Token {
	for i := 0; i < len(s.Source); i++ {
		s.Start = s.Current
		if !s.isAtEnd() {
			s.scanToken()
		}
	}

	s.Tokens = append(s.Tokens, token.Token{TokenType: token.EOF, Lexeme: "", Literal: nil, Line: s.Line})
	return s.Tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(token.LEFT_PAREN)
	case ')':
		s.addToken(token.RIGHT_PAREN)
	case '{':
		s.addToken(token.LEFT_BRACE)
	case '}':
		s.addToken(token.RIGHT_BRACE)
	case ',':
		s.addToken(token.COMMA)
	case '.':
		s.addToken(token.DOT)
	case '-':
		s.addToken(token.MINUS)
	case '+':
		s.addToken(token.PLUS)
	case ';':
		s.addToken(token.SEMICOLON)
	case '*':
		s.addToken(token.STAR)
	case '!':
		var ntt token.TType
		if s.match('=') {
			ntt = token.BANG_EQUAL
		} else {
			ntt = token.BANG
		}
		s.addToken(ntt)
	case '=':
		var ntt token.TType
		if s.match('=') {
			ntt = token.EQUAL_EQUAL
		} else {
			ntt = token.EQUAL
		}
		s.addToken(ntt)
	case '<':
		var ntt token.TType
		if s.match('=') {
			ntt = token.LESS_EQUAL
		} else {
			ntt = token.LESS
		}
		s.addToken(ntt)

	case '>':
		var ntt token.TType
		if s.match('=') {
			ntt = token.GREATER_EQUAL
		} else {
			ntt = token.GREATER
		}
		s.addToken(ntt)
	default:
		s.Errors = append(s.Errors, Error{Source: s.Source[s.Start:s.Current], Line: s.Line, Start: s.Start, Current: s.Current})
		fmt.Println("Error at line: ", s.Line, s.Source[s.Start:s.Current])
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.Current >= len(s.Source)
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	if string(s.Source[s.Current]) != string(expected) {
		return false
	}
	s.Current += 1
	return true
}

// TODO: returning a byte here is OK for now, but if we need to allow multi-type utf chars, it should be a rune.
func (s *Scanner) advance() byte {
	b := s.Source[s.Current]
	s.Current += 1
	return b
}

func (s *Scanner) addToken(tok token.TType) {
	s.addTokenWithObj(tok, nil)
}

func (s *Scanner) addTokenWithObj(tok token.TType, obj interface{}) {
	text := s.Source[s.Start:s.Current]
	s.Tokens = append(s.Tokens, token.MakeToken(tok, text, nil, s.Line))
}
