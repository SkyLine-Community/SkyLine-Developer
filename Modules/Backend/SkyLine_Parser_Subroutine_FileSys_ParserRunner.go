package SkyLine_Backend

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func RunAndParseNoFile(data string) error {
	parser := New_Parser(LexNew(data))
	program := parser.ParseProgram()
	if len(parser.ParserErrors()) > 0 {
		return errors.New(parser.ParserErrors()[0])
	}
	Env := NewEnvironment()
	result := Eval(program, Env)
	if _, ok := result.(*Nil); ok {
		return nil
	}
	defer func() {
		if xy := recover(); xy != nil {
			if strings.Contains(fmt.Sprint(xy), "invalid memory address or nil pointer dereference") {
				if *ErrorsTrace {
					fmt.Println(Map_Parser[ERROR_FILE_INTEGRITY_IS_EMPTY]("FILE_SECTOR_1:WAS_NOT_FILE_USED_FLAG_(-e) | data when using flag -e was empty, interpreter will NOT run"))
				}
			}
		}
	}()
	_, x := io.WriteString(os.Stdout, result.Inspect()+"\n")
	return x
}
