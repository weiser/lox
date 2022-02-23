package resolver

import (
	"github.com/golang-collections/collections/stack"
	"github.com/weiser/lox/expr"
	"github.com/weiser/lox/interpreter"
)

type Resolver struct {
	Interpreter interpreter.Interpreter
	Scopes      stack.Stack
}

func (r *Resolver) VisitExpr(e *expr.Expr) interface{}         { return nil }
func (r *Resolver) VisitAssign(e *expr.Assign) interface{}     { return nil }
func (r *Resolver) VisitBinary(e *expr.Binary) interface{}     { return nil }
func (r *Resolver) VisitCall(e *expr.Call) interface{}         { return nil }
func (r *Resolver) VisitGrouping(e *expr.Grouping) interface{} { return nil }
func (r *Resolver) VisitLiteral(e *expr.Literal) interface{}   { return nil }
func (r *Resolver) VisitLogical(e *expr.Logical) interface{}   { return nil }
func (r *Resolver) VisitUnary(e *expr.Unary) interface{}       { return nil }
func (r *Resolver) VisitVariable(e *expr.Variable) interface{} { return nil }

func (r *Resolver) VisitStmt(e *expr.Stmt) interface{}             { return nil }
func (r *Resolver) VisitExpression(e *expr.Expression) interface{} { return nil }
func (r *Resolver) VisitFunction(e *expr.Function) interface{}     { return nil }
func (r *Resolver) VisitIf(e *expr.If) interface{}                 { return nil }
func (r *Resolver) VisitPrint(e *expr.Print) interface{}           { return nil }
func (r *Resolver) VisitWhile(e *expr.While) interface{}           { return nil }
func (r *Resolver) VisitVar(e *expr.Var) interface{}               { return nil }
func (r *Resolver) VisitReturn(e *expr.Return) interface{}         { return nil }

func (r *Resolver) VisitBlock(block *expr.Block) interface{} {
	r.beginScope()
	r.resolveStatements(block.Statements)
	r.endScope()
	return nil
}

func (r *Resolver) resolveStatements(stmts []expr.StmtInterface) {
	for _, s := range stmts {
		r.resolveStatement(s)
	}
}

func (r *Resolver) resolveStatement(stmt expr.StmtInterface) {
	stmt.Accept(r)
}

func (r *Resolver) resolveExpression(exp expr.ExprInterface) {
	exp.Accept(r)
}

func (r *Resolver) beginScope() {
	r.Scopes.Push(map[string]bool{})
}

func (r *Resolver) endScope() {
	r.Scopes.Pop()
}
