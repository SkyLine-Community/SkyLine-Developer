package SkyLine_Backend

import (
	"bufio"
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

func RunAndParseFiles(filename string) error {
	FileCurrent.New(filename)
	f, x := os.Open(filename)
	if x != nil {
		defer func() {
			if x := recover(); x != nil {
				fmt.Println(x)
			}
		}()
		fmt.Println(Map_Parser[ERROR_FILE_INTEGRITY_DOES_NOT_EXIST](filename))
		os.Exit(0)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var line []string
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	if line == nil {
		fmt.Println(Map_Parser[ERROR_FILE_INTEGRITY_IS_EMPTY](filename))
	}
	data, x := ioutil.ReadFile(filename)
	if x != nil {
		fmt.Println(Map_Parser[ERROR_FILE_INPUT_OUTPUT_BUFFER_FAILED](filename, fmt.Sprint(x)))
		os.Exit(1)
	}
	parser := New_Parser(LexNew(string(data)))
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
					fmt.Println(Map_Parser[ERROR_FILE_INTEGRITY_IS_EMPTY](filename))
				}
			}
		}
	}()
	_, x = io.WriteString(os.Stdout, result.Inspect()+"\n")
	return x
}
