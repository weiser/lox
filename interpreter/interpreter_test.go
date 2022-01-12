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

	unary2 := &expr.Unary{Operator: token.MakeToken(token.BANG, "!", nil, 1), Right: &expr.Literal{Value: true}}
	actual2 := i.VisitUnary(unary2)
	if actual2 != false {
		t.Errorf("Expected false, got %v", actual2)
	}
}

func TestVisitBinary(t *testing.T) {
	i := &Interpreter{}
	bin1 := &expr.Binary{Operator: token.MakeToken(token.PLUS, "+", nil, 1), Right: &expr.Literal{Value: 5}, Left: &expr.Literal{Value: 6}}

	actual1 := i.VisitBinary(bin1)
	if actual1 != 11.0 {
		t.Errorf("Expected 11, got %v", actual1)
	}

	bin2 := &expr.Binary{Operator: token.MakeToken(token.PLUS, "+", nil, 1), Right: &expr.Literal{Value: "i"}, Left: &expr.Literal{Value: "h"}}

	actual2 := i.VisitBinary(bin2)
	if actual2 != "hi" {
		t.Errorf("Expected hi, got %v", actual2)
	}

	bin3 := &expr.Binary{Operator: token.MakeToken(token.GREATER, ">", nil, 1), Right: &expr.Literal{Value: 5}, Left: &expr.Literal{Value: 6}}

	actual3 := i.VisitBinary(bin3)
	if actual3 != true {
		t.Errorf("Expected true, got %v", actual3)
	}

	bin4 := &expr.Binary{Operator: token.MakeToken(token.EQUAL_EQUAL, "==", nil, 1), Right: &expr.Literal{Value: 5}, Left: &expr.Literal{Value: 5}}

	actual4 := i.VisitBinary(bin4)
	if actual4 != true {
		t.Errorf("Expected true, got %v", actual4)
	}

	bin5 := &expr.Binary{Operator: token.MakeToken(token.BANG_EQUAL, "!=", nil, 1), Right: &expr.Literal{Value: 5}, Left: &expr.Literal{Value: 5}}

	actual5 := i.VisitBinary(bin5)
	if actual5 != false {
		t.Errorf("Expected false, got %v", actual5)
	}
}
