package SkyLine_Backend

import "strconv"

func evalUnquoteCalls(quoted Node, Env *Environment_of_environment) Node {
	modifier := func(node Node) Node {
		call, ok := node.(*CallExpression)
		if !ok || call.Function.TokenLiteral() != FuncNameUnquote || len(call.Arguments) != 1 {
			return node
		}

		unquoted := Eval(call.Arguments[0], Env)
		return convertObjectToASTNode(unquoted)
	}
	return Modify(quoted, modifier)
}

func convertObjectToASTNode(obj Object) Node {
	switch obj := obj.(type) {
	case *Integer:
		base := 10
		t := Token{
			Token_Type: INT,
			Literal:    strconv.FormatInt(obj.Value, base),
		}
		return &IntegerLiteral{Token: t, Value: obj.Value}
	case *Boolean_Object:
		var t Token
		if obj.Value {
			t = Token{Token_Type: TRUE, Literal: "true"}
		} else {
			t = Token{Token_Type: FALSE, Literal: "false"}
		}
		return &Boolean_AST{Token: t, Value: obj.Value}
	case *Quote:
		return obj.Node
	default:
		return nil
	}
}

func ExpandMacros(program Node, Env *Environment_of_environment) Node {
	modifier := func(node Node) Node {
		call, ok := node.(*CallExpression)
		if !ok {
			return node
		}

		macro, ok := isMacroCall(call, Env)
		if !ok {
			return node
		}

		args := quoteArgs(call)
		evalEnv := extendMacroEnv(macro, args)

		quote, ok := Eval(macro.Body, evalEnv).(*Quote)
		if !ok {
			panic("we only support returning AST-nodes from macros")
		}

		return quote.Node
	}

	return Modify(program, modifier)
}
