package SkyLine_Backend

import (
	"fmt"
	"os"
)

func (e *Environment_of_environment) Set(name string, val Object) Object {
	current := e.Store[name]
	if current != nil && e.ROM[name] {
		fmt.Println("Attempting to set a value for a constant, mathematically and computationally it is illegal to change a constant")
		os.Exit(0)
	}
	e.Store[name] = val
	return val
}
