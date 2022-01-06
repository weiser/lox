package playground

import (
	"fmt"
	"strings"

	"github.com/weiser/lox/expr"
	"github.com/weiser/lox/token"
)

type AstPrinter struct {
	E  expr.ExprInterface
	Sb strings.Builder
}

func (a *AstPrinter) Print() string {
	a.E.Accept(a)
	return a.Sb.String()
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
		a.Sb.WriteString("nil")
	} else {
		a.Sb.WriteString(fmt.Sprintf("%v", e.Value))
	}

}
func (a *AstPrinter) VisitUnary(e *expr.Unary) {
	a.parenthesize(e.Operator.Lexeme, e.Right)
}
func (a *AstPrinter) parenthesize(lexeme string, rest ...expr.ExprInterface) {
	a.Sb.WriteString("(")
	a.Sb.WriteString(lexeme)
	for _, expr := range rest {
		a.Sb.WriteString(" ")
		expr.Accept(a)

	}
	a.Sb.WriteString(")")
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
	astp := AstPrinter{E: expr, Sb: strings.Builder{}}

	fmt.Println(astp.Print())

}
