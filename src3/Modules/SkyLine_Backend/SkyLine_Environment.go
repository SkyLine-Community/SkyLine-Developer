package SkyLine

import (
	"fmt"
	"os"
)

func NewEnvironment() Environment {
	return &Environment_of_environment{
		Store: make(map[string]Object),
		Outer: nil,
		ROM:   make(map[string]bool),
	}
}

func (e *Environment_of_environment) Get(name string) (Object, bool) {
	obj, exists := e.Store[name]
	if !exists && e.Outer != nil {
		obj, exists = e.Outer.Get(name)
	}
	return obj, exists
}

func (e *Environment_of_environment) Set(name string, val Object) Object {
	current := e.Store[name]
	if current != nil && e.ROM[name] {
		fmt.Println("Attempting to set a value for a constant, mathematically and computationally it is illegal to change a constant")
		os.Exit(0)
	}
	e.Store[name] = val
	return val
}

func NewEnclosedEnvironment(outer Environment) Environment {
	return &Environment_of_environment{
		Store: make(map[string]Object),
		Outer: outer,
	}
}

func (e *Environment_of_environment) SetConstants(name string, value Object) Object {
	e.Store[name] = value
	e.ROM[name] = true
	return value
}
