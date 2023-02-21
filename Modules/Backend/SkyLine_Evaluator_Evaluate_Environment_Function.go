package SkyLine_Backend

import (
	"fmt"
	"os"
)

func extendFunctionEnv(fn *Function, args []Object) *Environment_of_environment {
	Env := NewEnclosedEnvironment(fn.Env)

	for i, param := range fn.Parameters {
		Env.Set(param.Value, args[i])
		defer func() {
			if x := recover(); x != nil {
				fmt.Println(ErrorSymBolMap[CODE_PARSE_FUNCTION_ARGUMENTS_NOT_ENOUGH_ERROR](fmt.Sprint(fn.Inspect()), fmt.Sprint(len(args)), fmt.Sprint(len(fn.Parameters))))
				os.Exit(0)
			}
		}()
	}
	return Env
}

func DefineMacros(program *Program, Env *Environment_of_environment) {
	defs := make([]int, 0)
	stmts := program.Statements

	for pos, stmt := range stmts {
		if isMacroDefinition(stmt) {
			addMacro(stmt, Env)
			defs = append(defs, pos)
		}
	}

	for i := len(defs) - 1; i >= 0; i-- {
		pos := defs[i]
		program.Statements = append(stmts[:pos], stmts[pos+1:]...)
	}
}
func extendMacroEnv(macro *Macro, args []*Quote) *Environment_of_environment {
	extended := NewEnclosedEnvironment(macro.Env)
	for i, param := range macro.Parameters {
		extended.Set(param.Value, args[i])
	}
	return extended
}
