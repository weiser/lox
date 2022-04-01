package interpreter

import (
	"errors"
	"fmt"
	"time"

	"github.com/weiser/lox/environment"
	"github.com/weiser/lox/expr"
	"github.com/weiser/lox/token"
)

type LoxClass struct {
	Name    string
	Methods map[string]LoxFunction
}

/* loxclass needs to implement loxcallable so that we can do stuff like:
```
class A {}
A();
```
to instantiate an instance of A

*/
func (lc LoxClass) Arity() int {
	return 0
}

func (lc LoxClass) Call(i *Interpreter, arguments []interface{}) (retVal interface{}) {
	instance := LoxInstance{Klass: lc, Fields: make(map[string]interface{})}
	return instance
}

type LoxInstance struct {
	Klass  LoxClass
	Fields map[string]interface{}
}

func (li LoxInstance) String() string {
	return li.Klass.Name + " instance"
}

func (li LoxInstance) Get(name token.Token) interface{} {
	if v, ok := li.Fields[name.Lexeme]; ok {
		return v
	}

	method, ok := li.Klass.FindMethod(name.Lexeme)
	if ok {
		return method
	}

	panic(fmt.Sprintf("%v: undefined property '%v'", name, name.Lexeme))
}

func (li LoxInstance) Set(name token.Token, value interface{}) {
	li.Fields[name.Lexeme] = value
}

func (lc LoxClass) String() string {
	return lc.Name
}

func (lc LoxClass) FindMethod(name string) (LoxFunction, bool) {
	v, ok := lc.Methods[name]
	return v, ok
}

type LoxFunction struct {
	Declaration expr.Function
	Closure     environment.Environment
}

func (lf LoxFunction) Arity() int {
	return len(lf.Declaration.Params)
}

func (lf LoxFunction) Call(i *Interpreter, arguments []interface{}) (retVal interface{}) {
	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(ErrReturn); !ok {
				panic(err)
			} else {
				// here is where we catch any return values from a return statement
				retVal = v.Value
			}
		}

	}()
	environment := environment.MakeEnvironment(&lf.Closure)
	for i, p := range lf.Declaration.Params {
		environment.Define(p.Lexeme, arguments[i])
	}
	i.ExecuteBlock(lf.Declaration.Body, environment)
	return retVal
}

type LoxCallable interface {
	Arity() int
	Call(i *Interpreter, arguments []interface{}) interface{}
}

type ErrBreak struct {
}

type ErrReturn struct {
	Value interface{}
}

func (e *ErrBreak) Error() string {
	return fmt.Sprintf("Break encountered")
}

type Interpreter struct {
	env    environment.Environment
	Locals map[interface{}]int
}

var Globals environment.Environment

type GlobalClock struct{}

func (gclock *GlobalClock) Arity() int { return 0 }
func (gclock *GlobalClock) Call(i *Interpreter, arguments []interface{}) interface{} {
	return time.Now().UnixMilli()
}
func (gclock *GlobalClock) String() string {
	return "<native fxn: global clock>"
}

func InitGlobals() environment.Environment {
	Globals = environment.MakeEnvironment(nil)
	Globals.Define("clock", GlobalClock{})

	return Globals
}

func MakeInterpreter() Interpreter {
	return Interpreter{env: InitGlobals(), Locals: make(map[interface{}]int)}
}

func (i *Interpreter) VisitLiteral(exp *expr.Literal) interface{} {
	if exp.Value == token.BREAK {
		panic(ErrBreak{})
	}
	return exp.Value
}

func (i *Interpreter) VisitGrouping(exp *expr.Grouping) interface{} {
	return i.Evaluate(exp.Expression)
}

func (i *Interpreter) VisitBinary(exp *expr.Binary) interface{} {
	left := i.Evaluate(exp.Left)
	right := i.Evaluate(exp.Right)

	switch exp.Operator.TokenType {
	case token.BANG_EQUAL:
		return !isEqual(left, right)
	case token.EQUAL_EQUAL:
		return isEqual(left, right)
	case token.GREATER:
		lv, lok := toFloat(left)
		rv, rok := toFloat(right)
		if rok == nil && lok == nil {
			return lv > rv
		}
		fmt.Println("Tried to VisitBinary.GREATER and failed: ", left, right)
	case token.GREATER_EQUAL:
		lv, lok := toFloat(left)
		rv, rok := toFloat(right)
		if rok == nil && lok == nil {
			return lv >= rv
		}
		fmt.Println("Tried to VisitBinary.GREATER_EQUAL and failed: ", left, right)
	case token.LESS:
		lv, lok := toFloat(left)
		rv, rok := toFloat(right)
		if rok == nil && lok == nil {
			return lv < rv
		}
		fmt.Println("Tried to VisitBinary.LESS and failed: ", left, right)
	case token.LESS_EQUAL:
		lv, lok := toFloat(left)
		rv, rok := toFloat(right)
		if rok == nil && lok == nil {
			return lv <= rv
		}
		fmt.Println("Tried to VisitBinary.LESS_EQUAL and failed: ", left, right)
	case token.MINUS:
		lv, lok := toFloat(left)
		rv, rok := toFloat(right)
		if rok == nil && lok == nil {
			return lv - rv
		}
		fmt.Println("Tried to VisitBinary.MINUS and failed: ", left, right)
	case token.SLASH:
		lv, lok := toFloat(left)
		rv, rok := toFloat(right)
		if rok == nil && lok == nil {
			return lv / rv
		}
		fmt.Println("Tried to VisitBinary.SLASH and failed: ", left, right)
	case token.STAR:
		lv, lok := toFloat(left)
		rv, rok := toFloat(right)
		if rok == nil && lok == nil {
			return lv * rv
		}
		fmt.Println("Tried to VisitBinary.STAR and failed: ", left, right)
	case token.PLUS:
		lv, lok := toFloat(left)
		rv, rok := toFloat(right)
		if lok == nil && rok == nil {
			return lv + rv
		}

		lv2, lok2 := left.(string)
		rv2, rok2 := right.(string)
		if rok2 && lok2 {
			return lv2 + rv2
		}
		fmt.Println("Tried to VisitBinary.PLUS and failed: ", left, right)
	}

	return nil
}

func (i *Interpreter) VisitCall(call *expr.Call) interface{} {
	callee := i.Evaluate(call.Callee)
	arguments := make([]interface{}, 0)
	for _, arg := range call.Arguments {
		arguments = append(arguments, i.Evaluate(arg))
	}

	fxn, ok := callee.(LoxCallable)
	if !ok {
		panic(errors.New(fmt.Sprintf("%v: Can only call functions and classes", call.Paren)))
	}
	if len(arguments) != fxn.Arity() {
		panic(errors.New(fmt.Sprintf("Expected %v arguments, got %v arguments", fxn.Arity(), len(arguments))))
	}
	return fxn.Call(i, arguments)
}

func (i *Interpreter) VisitGet(get *expr.Get) interface{} {
	obj := i.Evaluate(get.Object)
	if li, ok := obj.(LoxInstance); ok {
		return li.Get(get.Name)
	}

	panic(fmt.Sprintf("%v: only instances have properties", get.Name))
}

func (i *Interpreter) VisitUnary(exp *expr.Unary) interface{} {
	right := i.Evaluate(exp.Right)

	switch exp.Operator.TokenType {
	case token.MINUS:
		v, err := toFloat(right)
		if err == nil {
			return -v
		}
		fmt.Println("Tried to VisitUnary.MINUS on ", exp, " and failed: ", err)
	case token.BANG:
		v, err := toTruthy(right)
		if err == nil {
			return !v
		}
		fmt.Println("Tried to VisitUnary.BANG on ", exp, " and failed: ", err)
	}

	return nil
}

func (i *Interpreter) VisitExpr(exp *expr.Expr) interface{} {
	// TODO
	return nil
}

func (i *Interpreter) VisitStmt(exp *expr.Stmt) interface{} {
	// TODO
	return nil
}

func (i *Interpreter) Evaluate(exp expr.ExprInterface) interface{} {
	value := exp.Accept(i)
	return value
}

func (i *Interpreter) VisitExpression(stmt *expr.Expression) interface{} {
	i.Evaluate(stmt.Expression)
	return nil
}

func (i *Interpreter) VisitFunction(fxn *expr.Function) interface{} {
	loxFxn := LoxFunction{Declaration: *fxn, Closure: i.env}
	i.env.Define(fxn.Name.Lexeme, loxFxn)
	return nil
}

func (i *Interpreter) VisitPrint(stmt *expr.Print) interface{} {
	value := i.Evaluate(stmt.Expression)
	fmt.Println(value)
	return nil
}

func (i *Interpreter) VisitVar(stmt *expr.Var) interface{} {
	var value interface{}
	value = nil
	if stmt.Initializer != nil {
		value = i.Evaluate(stmt.Initializer)
	}

	i.env.Define(stmt.Name.Lexeme, value)
	return nil
}

func (i *Interpreter) VisitReturn(ret *expr.Return) interface{} {
	if ret.Value != nil {
		value := i.Evaluate(ret.Value)
		panic(ErrReturn{Value: value})
	}
	panic(ErrReturn{})
}

func (i *Interpreter) VisitVariable(exp *expr.Variable) interface{} {
	i.LookupVariable(exp.Name, exp)
	v, err := i.env.Get(exp.Name.Lexeme)
	if err == nil {
		return v
	}
	panic(err)
}

func (i *Interpreter) LookupVariable(name token.Token, exp *expr.Variable) (interface{}, error) {
	distance := i.Locals[exp]

	if distance != 0 {
		return i.env.GetAt(distance, name.Lexeme), nil
	} else {
		return Globals.Get(name.Lexeme)
	}
}

func (i *Interpreter) VisitAssign(exp *expr.Assign) interface{} {
	value := i.Evaluate(exp.Value)

	distance := i.Locals[exp]
	if distance != 0 {
		i.env.AssignAt(distance, exp.Name, value)
	} else {
		Globals.Assign(exp.Name.Lexeme, value)
	}

	return value
}

func (i *Interpreter) EvaluateStmt(stmt expr.StmtInterface) interface{} {
	value := stmt.Accept(i)
	return value
}

func (i *Interpreter) Interpret(stmts []expr.StmtInterface) {
	for _, stmt := range stmts {
		i.Execute(stmt)
	}
}

func (i *Interpreter) Execute(stmt expr.StmtInterface) {
	stmt.Accept(i)
}

func (i *Interpreter) Resolve(exp expr.ExprInterface, depth int) {
	i.Locals[&exp] = depth
}

func (i *Interpreter) VisitBlock(block *expr.Block) interface{} {
	i.ExecuteBlock(block.Statements, environment.MakeEnvironment(&i.env))
	return nil
}

func (i *Interpreter) VisitClass(class *expr.Class) interface{} {
	i.env.Define(class.Name.Lexeme, nil)

	methods := make(map[string]LoxFunction)
	for _, method := range class.Methods {
		decl, ok := method.(*expr.Function)
		if ok {
			function := LoxFunction{Declaration: *decl, Closure: i.env}
			methods[decl.Name.Lexeme] = function
		}
	}

	klass := LoxClass{Name: class.Name.Lexeme, Methods: methods}
	i.env.Assign(klass.Name, klass)
	return nil
}

func (i *Interpreter) ExecuteBlock(stmts []expr.StmtInterface, env environment.Environment) {
	i2 := Interpreter{env: env}
	for _, stmt := range stmts {
		(&i2).Execute(stmt)
	}
}

func (i *Interpreter) VisitIf(ifStmt *expr.If) interface{} {
	if v, ok := toTruthy(i.Evaluate(ifStmt.Condition)); ok == nil && v {
		i.Execute(ifStmt.ThenBranch)
	} else {
		if ifStmt.ElseBranch != nil {
			i.Execute(ifStmt.ElseBranch)
		}
	}

	return nil
}

func (i *Interpreter) VisitLogical(logical *expr.Logical) interface{} {
	left := i.Evaluate(logical.Left)
	if logical.Operator.TokenType == token.OR {
		if v, ok := toTruthy(left); ok != nil && v {
			return left
		}
	} else {
		if v, ok := toTruthy(left); ok != nil && !v {
			return left
		}
	}
	return i.Evaluate(logical.Right)
}

func (i *Interpreter) VisitSet(set *expr.Set) interface{} {
	object := i.Evaluate(set.Object)
	li, ok := object.(LoxInstance)
	if !ok {
		panic(fmt.Sprintf("%v: only instances have fields", set.Name))
	}
	value := i.Evaluate(set.Value)
	li.Set(set.Name, value)
	return value

}

func (i *Interpreter) VisitWhile(stmt *expr.While) interface{} {
	defer func() {
		if err := recover(); err != nil {
			if _, ok := err.(ErrBreak); !ok {
				panic(err)
			}
		}
	}()
	v, _ := toTruthy(i.Evaluate(stmt.Condition))
	// if a stmt is a break, stop looping
	for v {
		i.Execute(stmt.Body)
		v, _ = toTruthy(i.Evaluate(stmt.Condition))
	}
	return nil
}

func toFloat(i interface{}) (float64, error) {
	switch v := i.(type) {
	case float64:
		return v, nil
	case int64:
		return float64(v), nil
	case int:
		return float64(v), nil
	}

	return 0.0, fmt.Errorf(" %v could not be parsed as float", i)
}

func toTruthy(i interface{}) (bool, error) {
	if i == nil {
		return false, nil
	}
	if v, ok := i.(bool); ok {
		return v, nil
	}
	// at this point, it isn't a boolean value, so anything else is truthy
	return true, nil
}

func isEqual(l interface{}, r interface{}) bool {
	return l == r
}
