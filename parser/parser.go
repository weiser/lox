package parser

import (
	"fmt"

	"github.com/weiser/lox/expr"
	"github.com/weiser/lox/token"
)

type ParserError struct {
	Token token.Token
	Msg   string
}

func (pe *ParserError) Error() string {
	return fmt.Sprintf("error on %v: %v", pe.Token, pe.Msg)
}

type Parser struct {
	Tokens  []token.Token
	Current int
}

func (p *Parser) Parse() (expr.ExprInterface, error) {
	var pe ParserError
	defer func() {
		if err := recover(); err != nil {
			v, ok := err.(ParserError)
			if ok {
				fmt.Println("got an error: ", v)
				pe = v
			} else {
				// any non-parsererror we will barf on
				panic(v)
			}
		}
	}()
	return p.Expression(), &pe
}

func (p *Parser) Expression() expr.ExprInterface {
	return p.Equality()
}

func (p *Parser) Equality() expr.ExprInterface {
	exp := p.Comparison()
	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.Comparison()
		exp = &expr.Binary{Left: exp, Operator: operator, Right: right}
	}
	return exp
}

func (p *Parser) Comparison() expr.ExprInterface {
	exp := p.Term()
	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right := p.Term()
		exp = &expr.Binary{Right: right, Operator: operator, Left: exp}
	}
	return exp
}

func (p *Parser) Term() expr.ExprInterface {
	exp := p.Factor()

	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right := p.Factor()
		exp = &expr.Binary{Left: exp, Operator: operator, Right: right}
	}
	return exp
}

func (p *Parser) Factor() expr.ExprInterface {
	exp := p.Unary()
	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right := p.Unary()
		exp = &expr.Binary{Left: exp, Operator: operator, Right: right}
	}
	return exp
}

func (p *Parser) Unary() expr.ExprInterface {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right := p.Unary()
		return &expr.Unary{Operator: operator, Right: right}
	}

	return p.Primary()
}

func (p *Parser) Primary() expr.ExprInterface {
	if p.match(token.FALSE) {
		return &expr.Literal{Value: false}
	}
	if p.match(token.TRUE) {
		return &expr.Literal{Value: true}
	}
	if p.match(token.NIL) {
		return &expr.Literal{Value: nil}
	}
	if p.match(token.NUMBER, token.STRING) {
		return &expr.Literal{Value: p.previous().Literal}
	}

	var e expr.ExprInterface
	if p.match(token.LEFT_PAREN) {
		exp := p.Expression()
		_, err := p.consume(token.RIGHT_PAREN, "Expect ')' after expression")
		if err == nil {
			e = &expr.Grouping{Expression: exp}
			return e
		} else {
			panic(err)
		}
	}
	err := MakeParserError(p.peek(), "expected expression")
	panic(err)

}

func (p *Parser) consume(tokenType token.TType, err string) (token.Token, error) {
	if p.checkType(tokenType) {
		return p.advance(), nil
	}

	return token.Token{}, MakeParserError(p.peek(), err)
}

func MakeParserError(tok token.Token, err string) error {
	return &ParserError{Token: tok, Msg: err}
}

func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().TokenType == token.SEMICOLON {
			return
		}

		switch p.peek().TokenType {
		case token.CLASS, token.FOR, token.FUN, token.IF, token.PRINT, token.RETURN, token.VAR, token.WHILE:
			return
		}
		p.advance()
	}
}

func (p *Parser) match(tokenTypes ...token.TType) bool {
	for _, typ := range tokenTypes {
		if p.checkType(typ) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) checkType(typ token.TType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().TokenType == typ
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.Current += 1
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == token.EOF
}

func (p *Parser) peek() token.Token {
	return p.Tokens[p.Current]
}

func (p *Parser) previous() token.Token {
	return p.Tokens[p.Current-1]
}
