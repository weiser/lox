package main

import (
	"fmt"
	"strings"

	"github.com/weiser/lox/expr"
	"github.com/weiser/lox/token"
)

type AstPrinter struct {
	e  expr.ExprInterface
	sb strings.Builder
}

func (a *AstPrinter) print() string {
	a.e.Accept(a)
	return a.sb.String()
}

func (a *AstPrinter) VisitExpr(e *expr.Expr) {
	e.Accept(a)
}

func (a *AstPrinter) VisitBinary(e *expr.Binary) {
	a.parenthesize(e.Operator.Lexeme, e.Left, e.Right)
}

func (a *AstPrinter) VisitGrouping(e *expr.Grouping) {
	a.parenthesize("group", e.Expression)
}

func (a *AstPrinter) VisitLiteral(e *expr.Literal) {
	if e.Value == nil {
		a.sb.WriteString("nil")
	} else {
		a.sb.WriteString(fmt.Sprintf("%v", e.Value))
	}

}
func (a *AstPrinter) VisitUnary(e *expr.Unary) {
	a.parenthesize(e.Operator.Lexeme, e.Right)
}
func (a *AstPrinter) parenthesize(lexeme string, rest ...expr.ExprInterface) {
	a.sb.WriteString("(")
	a.sb.WriteString(lexeme)
	for _, expr := range rest {
		a.sb.WriteString(" ")
		expr.Accept(a)

	}
	a.sb.WriteString(")")
}

func main() {
	expr := &expr.Binary{
		Left: &expr.Unary{
			Operator: token.MakeToken(token.MINUS, "-", nil, 1),
			Right:    &expr.Literal{Value: 123},
		},
		Operator: token.MakeToken(token.STAR, "*", nil, 1),
		Right: &expr.Grouping{
			Expression: &expr.Literal{Value: 45.67},
		},
	}
	astp := AstPrinter{e: expr, sb: strings.Builder{}}

	fmt.Println(astp.print())

}
