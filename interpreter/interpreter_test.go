package interpreter

import (
	"testing"

	"github.com/weiser/lox/expr"
)

func TestVisitLiteralExpr(t *testing.T) {
	i := &Interpreter{}

	literal := expr.Literal{Value: 5}

	actual := i.visitLiteralExpr(literal)
	if actual != 5 {
		t.Errorf("Expected 5, got %v", actual)
	}
}
