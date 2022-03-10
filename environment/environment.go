package environment

import (
	"fmt"

	"github.com/weiser/lox/token"
)

type Environment struct {
	Values    map[string]interface{}
	Enclosing *Environment
}

func MakeEnvironment(enclosing *Environment) Environment {
	return Environment{Values: make(map[string]interface{}), Enclosing: enclosing}
}

func (e *Environment) Define(name string, value interface{}) {
	e.Values[name] = value
}

func (e *Environment) Get(name string) (interface{}, error) {
	if value, ok := e.Values[name]; ok {
		return value, nil
	}

	if e.Enclosing != nil {
		return e.Enclosing.Get(name)
	}

	return nil, fmt.Errorf("variable %v is not defined", name)
}

func (e *Environment) GetAt(distance int, name string) interface{} {
	return e.Ancestor(distance).Values[name]
}

func (e *Environment) Ancestor(distance int) *Environment {
	curEnv := e
	for i := 0; i < distance; i += 1 {
		curEnv = curEnv.Enclosing
	}

	return curEnv
}

func (e *Environment) Assign(name string, value interface{}) (interface{}, error) {
	if _, ok := e.Values[name]; ok {
		e.Values[name] = value
		return value, nil
	}

	if e.Enclosing != nil {
		e.Enclosing.Assign(name, value)
		return value, nil
	}

	return nil, fmt.Errorf("variable %v is not defined", name)
}

func (e *Environment) AssignAt(distance int, name token.Token, value interface{}) {
	e.Ancestor(distance).Values[name.Lexeme] = value
}
