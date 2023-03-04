package SkyLine_Backend

import (
	"fmt"
	"os"
)

func Eval(node Node, Env *Environment_of_environment) Object {
	switch node := node.(type) {
	case *Program:
		return evalProgram(node, Env)
	case *ExpressionStatement:
		return Eval(node.Expression, Env)
	case *ReturnStatement:
		value := Eval(node.ReturnValue, Env)
		if isError(value) {
			return value
		}
		return &ReturnValue{Value: value}
	case *Constant:
		value := Eval(node.Value, Env)
		if isError(value) {
			return value
		}
		Env.Set(node.Name.Value, value)
		return value
	case *ObjectCallExpression:
		res := Evaluate_ObjectCall(node, Env)
		if isError(res) {
			fmt.Fprintf(os.Stderr, "Error calling method: %s", res.Inspect())
		}
		return res
	case *BlockStatement:
		return evalBlockStatement(node, Env)
	case *Switch:
		return EvalSwitch(node, Env)
	case *LetStatement:
		value := Eval(node.Value, Env)
		if isError(value) {
			return value
		}
		Env.Set(node.Name.Value, value)
	case *AssignmentStatement:
		return EvalAssignModify(node, Env)
	case *WhileCondition:
		return EvalWhileCondition(node, Env)
	// Expressions
	case *IntegerLiteral:
		return &Integer{Value: node.Value}

	case *FloatLiteral:
		return &Float{Value: node.Value}

	case *Boolean_AST:
		return nativeBoolToBooleanObject(node.Value)

	case *PrefixExpression:
		right := Eval(node.Right, Env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

	case *InfixExpression:
		left := Eval(node.Left, Env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, Env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)
	case *ConditionalExpression:
		return evalIfExpression(node, Env)
	case *Register:
		value := Eval(node.ImportValue, Env)
		if isError(value) {
			return value
		}
		return EvalRegisterCall(value)
	case *Ident:
		return evalIdent(node, Env)

	case *FunctionLiteral:
		return &Function{
			Parameters: node.Parameters,
			Body:       node.Body,
			Env:        Env,
		}
	case *CallExpression:
		if node.Function.TokenLiteral() == FuncNameQuote {
			return quote(node.Arguments[0], Env)
		}

		function := Eval(node.Function, Env)
		if isError(function) {
			return function
		}

		args := evalExpressions(node.Arguments, Env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return applyFunction(Env, function, args)

	case *StringLiteral:
		return &String{Value: node.Value}

	case *ArrayLiteral:
		elems := evalExpressions(node.Elements, Env)
		if len(elems) == 1 && isError(elems[0]) {
			return elems[0]
		}
		return &Array{Elements: elems}

	case *IndexExpression:
		left := Eval(node.Left, Env)
		if isError(left) {
			return left
		}
		index := Eval(node.Index, Env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)

	case *HashLiteral:
		return evalHashLiteral(node, Env)
	}

	return nil
}
