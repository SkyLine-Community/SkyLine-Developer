package SkyLine_Backend

import "fmt"

func Println(args ...Object) Object {
	for _, k := range args {
		fmt.Println(k.Inspect())
	}
	return NilValue
}

func Print(args ...Object) Object {
	for _, k := range args {
		fmt.Print(k.Inspect())
	}
	return NilValue
}

func Sprint(args ...Object) Object {
	if len(args) == 0 {
		return NewError("Sprint takes at least one positional argunent")
	}
	var returnedout string
	for _, k := range args {
		returnedout += fmt.Sprint(k.Inspect())
	}
	return &String{Value: returnedout}
}
