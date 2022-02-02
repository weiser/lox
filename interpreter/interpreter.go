package interpreter

import (
	"fmt"

	"github.com/weiser/lox/environment"
	"github.com/weiser/lox/expr"
	"github.com/weiser/lox/token"
)

type Interpreter struct {
	env environment.Environment
}

func MakeInterpreter() Interpreter {
	return Interpreter{env: environment.MakeEnvironment(nil)}
}

func (i *Interpreter) VisitLiteral(exp *expr.Literal) interface{} {
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

func (i *Interpreter) VisitVariable(exp *expr.Variable) interface{} {
	v, err := i.env.Get(exp.Name.Lexeme)
	if err == nil {
		return v
	}
	panic(err)
}

func (i *Interpreter) VisitAssign(exp *expr.Assign) interface{} {
	value := i.Evaluate(exp.Value)
	_, err := i.env.Assign(exp.Name.Lexeme, value)
	if err != nil {
		panic(err)
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

// TODO text.lox isn't setting environmnents correctly
func (i *Interpreter) VisitBlock(block *expr.Block) interface{} {
	i.ExecuteBlock(block.Statements, environment.MakeEnvironment(&i.env))
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

func (i *Interpreter) VisitWhile(stmt *expr.While) interface{} {
	v, _ := toTruthy(i.Evaluate(stmt.Condition))
	// todo pg 149 for implementing loops+break:
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
