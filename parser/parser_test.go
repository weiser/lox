package parser

import (
	"testing"

	"github.com/weiser/lox/expr"
	"github.com/weiser/lox/scanner"
)

func TestIfStmt(t *testing.T) {
	scanner := scanner.MakeScanner(`if (1>0) { print 1;}`)
	toks := scanner.ScanTokens()
	p := Parser{Tokens: toks}
	stmts, _ := p.Parse()

	if _, ok := stmts[0].(*expr.If); !ok {
		t.Errorf("expected an If statement, got a %v", stmts[0])
	}
}

func TestLogicalExpr(t *testing.T) {
	scanner := scanner.MakeScanner(`var a = null or 1;`)
	toks := scanner.ScanTokens()
	p := Parser{Tokens: toks}
	stmts, _ := p.Parse()

	v, ok := stmts[0].(*expr.Var)
	if !ok {
		t.Errorf("expected an asignment statement, got a %v", stmts[0])
	}
	if _, ok := v.Initializer.(*expr.Logical); !ok {
		t.Errorf("Expected a logical expro, got a, %v", v)
	}
}

func TestWhileStmt(t *testing.T) {
	scanner := scanner.MakeScanner(`while (true) {
		print 1;
	}`)
	toks := scanner.ScanTokens()
	p := Parser{Tokens: toks}
	stmts, _ := p.Parse()

	_, ok := stmts[0].(*expr.While)
	if !ok {
		t.Errorf("expected a while statement, got a %v", stmts[0])
	}
}

func TestForStmt(t *testing.T) {
	scanner := scanner.MakeScanner(`for (var a= 1; a < 10; a = a+1) {
		print 1;
	}`)
	toks := scanner.ScanTokens()
	p := Parser{Tokens: toks}
	stmts, _ := p.Parse()

	block, ok := stmts[0].(*expr.Block)
	if !ok {
		t.Errorf("expected a Block statement, got a %v", stmts[0])
	}

	if _, ok := block.Statements[0].(*expr.Var); !ok {
		t.Errorf("expected a Variableexpression statement, got a %v", block.Statements[0])
	}
	if _, ok := block.Statements[1].(*expr.While); !ok {
		t.Errorf("expected a while statement, got a %v", block.Statements[1])
	}
}

func TestClassStmt(t *testing.T) {
	scanner := scanner.MakeScanner(`class Test { t() {return 1;}}`)
	toks := scanner.ScanTokens()
	p := Parser{Tokens: toks}
	stmts, _ := p.Parse()

	_, ok := stmts[0].(*expr.Class)
	if !ok {
		t.Errorf("expected a Class statement, got a %v", stmts[0])
	}
}
