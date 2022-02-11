package interpreter

import (
	"testing"

	"github.com/weiser/lox/expr"
	"github.com/weiser/lox/parser"
	"github.com/weiser/lox/scanner"
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

func TestExprStmt(t *testing.T) {
	scanner := scanner.MakeScanner("1+2;")
	parser := parser.Parser{Tokens: scanner.ScanTokens()}
	stmts, err := parser.Parse()
	if err != nil {
		t.Errorf("didn't parse, %v", err)
	}
	i := &Interpreter{}

	i.Interpret(stmts)
}

func TestPrintStmt(t *testing.T) {
	scanner := scanner.MakeScanner("print 1+2;")
	parser := parser.Parser{Tokens: scanner.ScanTokens()}
	stmts, err := parser.Parse()
	if err != nil {
		t.Errorf("didn't parse, %v", err)
	}
	i := &Interpreter{}

	i.Interpret(stmts)
}

func TestAssignStmt(t *testing.T) {
	scanner := scanner.MakeScanner("var a = \"global\"; print a;")
	parser := parser.Parser{Tokens: scanner.ScanTokens()}
	stmts, err := parser.Parse()
	if err != nil {
		t.Errorf("didn't parse, %v", err)
	}
	i := MakeInterpreter()

	(&i).Interpret(stmts)
	o, _ := (&i).env.Get("a")
	a := o.(string)
	if a != "global" {
		t.Errorf("expected a = 'global', ggot a = '%v'", a)
	}

}

func TestBlockStmt(t *testing.T) {
	scanner := scanner.MakeScanner(`
	var a = 1;
	var b = 2;
	var c = 0;
	{
		var a = 3;
		c = a + b;
	}
	`)
	parser := parser.Parser{Tokens: scanner.ScanTokens()}
	stmts, err := parser.Parse()
	if err != nil {
		t.Errorf("didn't parse, %v", err)
	}
	i := MakeInterpreter()

	(&i).Interpret(stmts)
	var a, b, c float64
	o, _ := (&i).env.Get("a")
	a = o.(float64)
	o, _ = (&i).env.Get("b")
	b = o.(float64)
	o, _ = (&i).env.Get("c")
	c = o.(float64)
	if a != 1 || b != 2 || c != 5 {
		t.Errorf("should have gotten a = 1, b = 2, c = 5. Got a = %v, b = %v, c = %v", a, b, c)
	}

}

func TestIfStmt(t *testing.T) {
	scanner := scanner.MakeScanner(`
	var a = 1;
	if (a > 0) {
	  a = 2;
	} else {
	  a = 3;
	}
	`)
	parser := parser.Parser{Tokens: scanner.ScanTokens()}
	stmts, err := parser.Parse()
	if err != nil {
		t.Errorf("didn't parse, %v", err)
	}
	i := MakeInterpreter()

	(&i).Interpret(stmts)
	var a float64
	o, _ := (&i).env.Get("a")
	a = o.(float64)
	if a != 2 {
		t.Errorf("expected a = 2, instead a = %v", a)
	}

}

func TestLogicalStmt(t *testing.T) {
	scanner := scanner.MakeScanner(`
	var a = nil or 1;	
	`)
	parser := parser.Parser{Tokens: scanner.ScanTokens()}
	stmts, err := parser.Parse()
	if err != nil {
		t.Errorf("didn't parse, %v", err)
	}
	i := MakeInterpreter()

	(&i).Interpret(stmts)
	var a float64
	o, _ := (&i).env.Get("a")
	a = o.(float64)
	if a != 1 {
		t.Errorf("expected a = 1, instead a = %v", a)
	}
}

func TestWhileStmt(t *testing.T) {
	scanner := scanner.MakeScanner(`
	var a = 0;
	while (a == 0) {
		a = 1;
	}	
	`)
	parser := parser.Parser{Tokens: scanner.ScanTokens()}
	stmts, err := parser.Parse()
	if err != nil {
		t.Errorf("didn't parse, %v", err)
	}
	i := MakeInterpreter()

	(&i).Interpret(stmts)
	var a float64
	o, _ := (&i).env.Get("a")
	a = o.(float64)
	if a != 1 {
		t.Errorf("expected a = 1, instead a = %v", a)
	}
}

func TestForStmt(t *testing.T) {
	scanner := scanner.MakeScanner(`
	var a = 0;
	var b = 0;
	for (; a < 10; a = a + 1) {
		b = a;
	}	
	`)
	parser := parser.Parser{Tokens: scanner.ScanTokens()}
	stmts, err := parser.Parse()
	if err != nil {
		t.Errorf("didn't parse, %v", err)
	}
	i := MakeInterpreter()

	(&i).Interpret(stmts)
	var a float64
	o, _ := (&i).env.Get("a")
	a = o.(float64)
	if a != 10 {
		t.Errorf("expected a = 9, instead a = %v", a)
	}
	var b float64
	o, _ = (&i).env.Get("b")
	b = o.(float64)
	if b != 9 {
		t.Errorf("expected b = 9, instead b = %v", b)
	}
}

func TestForStmtWithBreak(t *testing.T) {
	scanner := scanner.MakeScanner(`
	var a = 0;
	var b = 1;
	for (; a < 10; a = a + 1) {
		break;
		b = a;
	}	
	`)
	parser := parser.Parser{Tokens: scanner.ScanTokens()}
	stmts, err := parser.Parse()
	if err != nil {
		t.Errorf("didn't parse, %v", err)
	}
	i := MakeInterpreter()

	(&i).Interpret(stmts)
	var a float64
	o, _ := (&i).env.Get("a")
	a = o.(float64)
	if a != 0 {
		t.Errorf("expected a = 0, instead a = %v", a)
	}
	var b float64
	o, _ = (&i).env.Get("b")
	b = o.(float64)
	if b != 1 {
		t.Errorf("expected b = 1, instead b = %v", b)
	}
}

func TestFunction(t *testing.T) {
	scanner := scanner.MakeScanner(`
	fun hi(a) {  print "hi, " + a + "!"; }
	hi("mom");
	`)
	parser := parser.Parser{Tokens: scanner.ScanTokens()}
	stmts, err := parser.Parse()
	if err != nil {
		t.Errorf("didn't parse, %v", err)
	}
	i := MakeInterpreter()

	(&i).Interpret(stmts)
}
