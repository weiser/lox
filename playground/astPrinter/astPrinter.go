package astPrinter

import (
	"fmt"
	"strings"

	"github.com/weiser/lox/expr"
)

type AstPrinter struct {
	E  expr.ExprInterface
	Sb strings.Builder
}

func (a *AstPrinter) Print() string {
	a.E.Accept(a)
	return a.Sb.String()
}

func (a *AstPrinter) VisitExpr(e *expr.Expr) interface{} {
	return e.Accept(a)
}

func (a *AstPrinter) VisitBinary(e *expr.Binary) interface{} {
	a.parenthesize(e.Operator.Lexeme, e.Left, e.Right)
	return nil
}

func (a *AstPrinter) VisitGrouping(e *expr.Grouping) interface{} {
	a.parenthesize("group", e.Expression)
	return nil
}

func (a *AstPrinter) VisitLiteral(e *expr.Literal) interface{} {
	if e.Value == nil {
		a.Sb.WriteString("nil")
	} else {
		a.Sb.WriteString(fmt.Sprintf("%v", e.Value))
	}
	return nil
}
func (a *AstPrinter) VisitUnary(e *expr.Unary) interface{} {
	a.parenthesize(e.Operator.Lexeme, e.Right)
	return nil
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
