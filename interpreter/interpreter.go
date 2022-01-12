package interpreter

import (
	"fmt"

	"github.com/weiser/lox/expr"
	"github.com/weiser/lox/token"
)

type Interpreter struct {
}

func (i *Interpreter) VisitLiteral(exp *expr.Literal) interface{} {
	return exp.Value
}

func (i *Interpreter) VisitGrouping(exp *expr.Grouping) interface{} {
	return i.evaluate(exp.Expression)
}

func (i *Interpreter) VisitBinary(exp *expr.Binary) interface{} {
	left := i.evaluate(exp.Left)
	right := i.evaluate(exp.Right)

	switch exp.Operator.TokenType {
	case token.BANG_EQUAL:
		return !isEqual(left, right)
	case token.EQUAL_EQUAL:
		return isEqual(left, right)
	case token.GREATER:
		lv, _ := toFloat(left)
		rv, _ := toFloat(right)
		return lv > rv
	case token.GREATER_EQUAL:
		lv, _ := toFloat(left)
		rv, _ := toFloat(right)
		return lv >= rv
	case token.LESS:
		lv, _ := toFloat(left)
		rv, _ := toFloat(right)
		return lv < rv
	case token.LESS_EQUAL:
		lv, _ := toFloat(left)
		rv, _ := toFloat(right)
		return lv <= rv
	case token.MINUS:
		lv, _ := toFloat(left)
		rv, _ := toFloat(right)
		return lv - rv
	case token.SLASH:
		lv, _ := toFloat(left)
		rv, _ := toFloat(right)
		return lv / rv
	case token.STAR:
		lv, _ := toFloat(left)
		rv, _ := toFloat(right)
		return lv * rv
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
	}

	return nil
}

func (i *Interpreter) VisitUnary(exp *expr.Unary) interface{} {
	right := i.evaluate(exp.Right)

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

func (i *Interpreter) evaluate(exp expr.ExprInterface) interface{} {
	return exp.Accept(i)
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
