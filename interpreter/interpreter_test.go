package interpreter

import (
	"testing"

	"github.com/weiser/lox/expr"
	"github.com/weiser/lox/token"
)

func TestVisitLiteralExpr(t *testing.T) {
	i := &Interpreter{}

	literal := &expr.Literal{Value: 5}

	actual := i.VisitLiteral(literal)
	if actual != 5 {
		t.Errorf("Expected 5, got %v", actual)
	}
}

func TestVisitUnary(t *testing.T) {
	i := &Interpreter{}

	unary := &expr.Unary{Operator: token.MakeToken(token.MINUS, "-", nil, 1), Right: &expr.Literal{Value: 5.0}}

	actual := i.VisitUnary(unary)
	if actual != -5.0 {
		t.Errorf("Expected -5, got %v", actual)
	}
}
