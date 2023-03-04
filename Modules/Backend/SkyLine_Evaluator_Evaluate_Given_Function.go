package SkyLine_Backend

func isMacroDefinition(node Statement) bool {
	letStmt, ok := node.(*LetStatement)
	if !ok {
		return false
	}

	_, ok = letStmt.Value.(*MacroLiteral)
	return ok
}

func addMacro(stmt Statement, Env *Environment_of_environment) {
	letStmt := stmt.(*LetStatement)
	macroLit := letStmt.Value.(*MacroLiteral)
	macro := &Macro{
		Parameters: macroLit.Parameters,
		Env:        Env,
		Body:       macroLit.Body,
	}
	Env.Set(letStmt.Name.Value, macro)
}

func isMacroCall(call *CallExpression, Env *Environment_of_environment) (macro *Macro, ok bool) {
	ident, ok := call.Function.(*Ident)
	if !ok {
		return nil, false
	}

	obj, ok := Env.Get(ident.Value)
	if !ok {
		return nil, false
	}

	macro, ok = obj.(*Macro)
	return macro, ok
}

func quoteArgs(call *CallExpression) []*Quote {
	args := make([]*Quote, 0, len(call.Arguments))
	for _, arg := range call.Arguments {
		args = append(args, &Quote{Node: arg})
	}
	return args
}
