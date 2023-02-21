package SkyLine_Backend

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadImportedCarrierFile(filename string) (bool, error) {
	data, x := ioutil.ReadFile(filename)
	if x != nil {
		log.Fatal(x)
	}
	parser := New_Parser(LexNew(string(data)))
	program := parser.ParseProgram()
	if len(parser.ParserErrors()) > 0 {
		return false, errors.New(parser.ParserErrors()[0])
	}

	Env := NewEnvironment()
	result := Eval(program, Env)
	if _, ok := result.(*Nil); ok {
		return false, nil
	}
	defer func() {
		if xy := recover(); xy != nil {
			if strings.Contains(fmt.Sprint(xy), "invalid memory address or nil pointer dereference") {
				if *ErrorsTrace {
					fmt.Println("SkyLine parser: no functions or variables loaded...")
				}
			}
		}
	}()
	_, x = io.WriteString(os.Stdout, result.Inspect()+"\n")
	if x != nil {
		return false, x
	}
	return true, nil
}
