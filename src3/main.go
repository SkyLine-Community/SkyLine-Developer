package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	Mod "main/Modules/SkyLine_Backend"
	Server "main/Modules/SkyLine_Server"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func init() {
	flag.Parse()
}

func main() {
	if *Mod.Server {
		Server.Server(false)
		time.Sleep(1 * time.Second)
	}
	if *Mod.Help {
		Mod.Banner()
		fmt.Println("Welcome to SkyLine the language for cyber security! Below you will find a \n Menu dedicated to running the interpreter with command line \n flags.")
		list := `
[----- Flag name -----]       |    [----- Flag description -----]
[--------------------------------------------------------------------]
--source                      | Load a source code file into SkyLine
--help                        | Display help
--bout                        | Enable SkyLine logo upon runtime
--trace                       | Will trace any attempted crashes during runtime

NoFlag = Run SkyLine's console
		`
		fmt.Println(list)
		os.Exit(0)
	}
	if *Mod.SourceFile != "" {
		if *Mod.Bnn {
			Mod.Banner()
		}
		if fileext := filepath.Ext(*Mod.SourceFile); fileext == ".csc" {
			if err := Run(*Mod.SourceFile); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(0)
			}
		} else {
			fmt.Println("Woah there buddy, sorry but that file type is not allowed, files must end in (.csc) not -> ", fileext)
		}
	} else {
		fmt.Println("\x1b[H\x1b[2J\x1b[3J")
		Mod.Banner()
		Mod.Start(os.Stdin, os.Stdout)
		return
	}
}

func Run(filename string) error {
	Mod.FileCurrent.New(filename)
	f, x := os.Open(filename)
	if x != nil {
		fmt.Println(Mod.ErrorSymBolMap[Mod.CODE_FILE_INTEGRITY_FILE_INVALID_FILE_DOES_NOT_EXIST](filename))
		os.Exit(0)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var line []string
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	if line == nil {
		fmt.Println(Mod.ErrorSymBolMap[Mod.CODE_FILE_INTEGRITY_FILE_INVALID_MUST_HAVE_CODE_OR_LOGIC_INSIDE_FILE_NULL](filename))
	}
	data, x := ioutil.ReadFile(filename)
	if x != nil {
		fmt.Println(Mod.ErrorSymBolMap[Mod.CODE_FILE_FAILED_USING_INPUT_OUTPUT_READER_AND_UTILITY_FILE_ISSUE](filename, fmt.Sprint(x)))
		os.Exit(1)
	}
	parser := Mod.New_Parser(Mod.LexNew(string(data)))
	program := parser.ParseProgram()
	if len(parser.ParserErrors()) > 0 {
		return errors.New(parser.ParserErrors()[0])
	}

	Env := Mod.NewEnvironment()
	result := Mod.Eval(program, Env)
	if _, ok := result.(*Mod.Nil); ok {
		return nil
	}
	defer func() {
		if xy := recover(); xy != nil {
			if strings.Contains(fmt.Sprint(xy), "invalid memory address or nil pointer dereference") {
				if *Mod.ErrorsTrace {
					fmt.Println(Mod.ErrorSymBolMap[Mod.CODE_NO_FUNCTIONS_OR_SYMBOLS_LOADED](filename))
				}
			}
		}
	}()
	_, x = io.WriteString(os.Stdout, result.Inspect()+"\n")

	return x
}
