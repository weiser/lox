package resolver

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
	"github.com/weiser/lox/expr"
	"github.com/weiser/lox/interpreter"
	"github.com/weiser/lox/token"
)

type Scope map[string]bool

type Stack struct {
	stack.Stack
}

type FunctionType int

const (
	NONE FunctionType = iota
	FUNCTION
)

type Resolver struct {
	Interpreter     interpreter.Interpreter
	Scopes          Stack
	CurrentFunction FunctionType
}

// Get's the ith item in the stack.  retains order of stack
func (s *Stack) Get(i int) interface{} {
	oldStack := stack.Stack{}
	if i == 0 {
		return s.Peek()
	}
	var ith interface{}
	for ind := 0; ind < i; ind += 1 {
		ith = s.Pop()
		oldStack.Push(ith)
	}
	for oldStack.Len() != 0 {
		s.Push(oldStack.Pop)
	}
	return ith
}

func (r *Resolver) VisitExpr(e *expr.Expr) interface{} { return nil }
func (r *Resolver) VisitAssign(e *expr.Assign) interface{} {
	r.resolveExpression(e.Value)
	r.resolveLocal(e, e.Name)
	return nil
}

func (r *Resolver) VisitBinary(e *expr.Binary) interface{} {
	r.resolveExpression(e.Right)
	r.resolveExpression(e.Left)
	return nil
}

func (r *Resolver) VisitCall(e *expr.Call) interface{} {
	r.resolveExpression(e.Callee)
	for _, arg := range e.Arguments {
		r.resolveExpression(arg)
	}
	return nil
}

func (r *Resolver) VisitGrouping(e *expr.Grouping) interface{} {
	r.resolveExpression(e.Expression)
	return nil
}

func (r *Resolver) VisitLiteral(e *expr.Literal) interface{} {
	return nil
}

func (r *Resolver) VisitLogical(e *expr.Logical) interface{} {
	r.resolveExpression(e.Right)
	r.resolveExpression(e.Left)
	return nil
}

func (r *Resolver) VisitUnary(e *expr.Unary) interface{} {
	r.resolveExpression(e.Right)
	return nil
}

func (r *Resolver) VisitVariable(e *expr.Variable) interface{} {
	if !(r.Scopes.Len() == 0) {
		scope, ok := r.Scopes.Peek().(Scope)
		if !ok {
			panic("scope wasn't seen in 'VisitVariable'")
		}
		if !scope[e.Name.Lexeme] {
			panic(fmt.Sprintln("Can't read local variable in its own initializer: %v", e.Name.Lexeme))
		}
	}
	r.resolveLocal(e, e.Name)
	return nil
}

func (r *Resolver) VisitStmt(e *expr.Stmt) interface{} { return nil }
func (r *Resolver) VisitExpression(e *expr.Expression) interface{} {
	r.resolveExpression(e.Expression)
	return nil
}
func (r *Resolver) VisitFunction(e *expr.Function) interface{} {
	r.declare(e.Name)
	r.define(e.Name)
	r.resolveFunction(e, FUNCTION)
	return nil
}
func (r *Resolver) VisitIf(e *expr.If) interface{} {
	r.resolveExpression(e.Condition)
	r.resolveStatement(e.ThenBranch)
	if e.ElseBranch != nil {
		r.resolveStatement(e.ElseBranch)
	}
	return nil
}
func (r *Resolver) VisitPrint(e *expr.Print) interface{} {
	r.resolveExpression(e.Expression)
	return nil
}
func (r *Resolver) VisitWhile(e *expr.While) interface{} {
	r.resolveExpression(e.Condition)
	r.resolveStatement(e.Body)
	return nil
}
func (r *Resolver) VisitVar(e *expr.Var) interface{} {
	r.declare(e.Name)
	if e.Initializer != nil {
		r.resolveExpression(e.Initializer)
	}
	r.define(e.Name)
	return nil
}
func (r *Resolver) VisitReturn(e *expr.Return) interface{} {
	if r.CurrentFunction == NONE {
		panic(fmt.Sprintf("'%v' cannot return from top level code", e.Keyword))
	}
	if e.Value != nil {
		r.resolveExpression(e.Value)
	}
	return nil
}

func (r *Resolver) VisitBlock(block *expr.Block) interface{} {
	r.beginScope()
	r.ResolveStatements(block.Statements)
	r.endScope()
	return nil
}

func (r *Resolver) VisitClass(class *expr.Class) interface{} {
	r.declare(class.Name)
	r.define(class.Name)
	return nil
}

func (r *Resolver) ResolveStatements(stmts []expr.StmtInterface) bool {
	successfullyResolved := true
	defer func() {
		if err := recover(); err != nil {
			successfullyResolved = false
		}
	}()

	for _, s := range stmts {
		r.resolveStatement(s)
	}

	return successfullyResolved
}

func (r *Resolver) resolveStatement(stmt expr.StmtInterface) {
	stmt.Accept(r)
}

func (r *Resolver) resolveExpression(exp expr.ExprInterface) {
	exp.Accept(r)
}

func (r *Resolver) resolveFunction(f *expr.Function, typ FunctionType) {
	enclosingType := r.CurrentFunction
	r.CurrentFunction = typ
	r.beginScope()
	for _, param := range f.Params {
		r.declare(param)
		r.define(param)
	}
	r.ResolveStatements(f.Body)
	r.endScope()
	r.CurrentFunction = enclosingType
}

func (r *Resolver) beginScope() {
	r.Scopes.Push(Scope{})
}

func (r *Resolver) endScope() {
	r.Scopes.Pop()
}

func (r *Resolver) declare(name token.Token) {
	if r.Scopes.Len() == 0 {
		return
	}

	scope, ok := r.Scopes.Peek().(Scope)
	if !ok {
		panic("scope wasn't valid")
	}
	if _, present := scope[name.Lexeme]; present {
		panic(fmt.Sprintf("variable '%v' already exists in scope", name.Lexeme))
	}
	scope[name.Lexeme] = false
}

func (r *Resolver) define(name token.Token) {
	if r.Scopes.Len() == 0 {
		return
	}
	scope, ok := r.Scopes.Peek().(Scope)
	if !ok {
		panic("scope not right type in 'define'")
	}
	scope[name.Lexeme] = true
}

func (r *Resolver) resolveLocal(e expr.ExprInterface, tok token.Token) {
	for i := r.Scopes.Len() - 1; i >= 0; i = i - 1 {
		scope, ok := r.Scopes.Get(i).(Scope)
		if ok && scope[tok.Lexeme] {
			r.Interpreter.Resolve(e, r.Scopes.Len()-1-i)
			return
		}
	}
}
