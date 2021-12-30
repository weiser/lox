package token

import "fmt"

type TType int

const (
	// single character tokens
	LEFT_PAREN TType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// one or two character tokens
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	//literals
	IDENTIFIER
	STRING
	NUMBER

	//keywords
	AND
	CLASS
	ELSE
	FALSE
	TRUE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	VAR
	WHILE

	EOF
)

type Token struct {
	TokenType TType
	Lexeme    string
	Literal   interface{}
	Line      int
}

func MakeToken(typ TType, lexeme string, literal interface{}, lineno int) Token {
	return Token{TokenType: typ, Lexeme: lexeme, Literal: literal, Line: lineno}
}

func (t Token) String() string {
	return fmt.Sprintf("token=%v lexeme=%v literal=%v line=%v", t.TokenType, t.Lexeme, t.Literal, t.Line)
}
