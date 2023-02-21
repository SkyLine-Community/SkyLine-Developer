package SkyLine_Backend

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func EvalSwitch(Switch *Switch, Env *Environment_of_environment) Object {
	object := Eval(Switch.Value, Env)
	for _, options := range Switch.Choices {
		if options.Def {
			continue
		}
		for _, values := range options.Expr {
			output := Eval(values, Env)
			if object.Type_Object() == output.Type_Object() && (object.Inspect() == output.Inspect()) {
				Block := evalBlockStatement(options.Block, Env)
				return Block
			}
		}
	}
	for _, opt := range Switch.Choices {
		if opt.Def {
			output := evalBlockStatement(opt.Block, Env)
			return output
		}
	}
	return nil
}

func Evaluate_ObjectCall(call *ObjectCallExpression, env *Environment_of_environment) Object {
	var meth *CallExpression
	obj := Eval(call.Object, env)
	if method, ok := call.Call.(*CallExpression); ok {
		args := evalExpressions(call.Call.(*CallExpression).Arguments, env)
		ret := obj.InvokeMethod(method.Function.String(), *env, args...)
		if ret != nil {
			return ret
		}
		attempts := []string{}
		attempts = append(attempts, strings.ToLower(string(obj.Type_Object())))
		attempts = append(attempts, "object")
		for _, prefix := range attempts {
			name := prefix + "." + method.Function.String()
			if fn, ok := env.Get(name); ok {
				extendEnv := extendFunctionEnv(fn.(*Function), args)
				extendEnv.Set("self", obj)
				evaluated := Eval(fn.(*Function).Body, extendEnv)
				obj = unwrapReturnValue(evaluated)
				return obj
			}
		}
	} else {
		meth = method
	}
	return NewError("Failed to invoke method: %s with method %s and object %s",
		call.Call.(*CallExpression).Function.String(),
		meth,
		obj,
	)
}

func quote(node Node, Env *Environment_of_environment) Object {
	node = evalUnquoteCalls(node, Env)
	return &Quote{Node: node}
}

func applyFunction(Env *Environment_of_environment, fn Object, args []Object) Object {
	switch fn := fn.(type) {
	case *Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *Builtin:
		return fn.Fn(Env, args...)
	default:
		// expecting crash if it is not a function type object | Macros are bugged
		defer func() {
			if x := recover(); x != nil {
				fmt.Print("\n\n Assuming this is a macro, there was something wrong. Macro does not exist")
				os.Exit(1)
			}
		}()
		return NewError(ErrorSymBolMap[CODE_PREFIX_PARSE_FUNCTION_INVALID_OR_UNFOUND_WITHIN_PARSER_AND_INTERPRETRR](string(fn.Inspect())))
	}
}

func unwrapReturnValue(obj Object) Object {
	if returnValue, ok := obj.(*ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func evalIndexExpression(left, index Object) Object {
	switch {
	case left.Type_Object() == ArrayType && index.Type_Object() == IntegerType:
		return evalArrayIndexExpression(left, index)
	case left.Type_Object() == HashType:
		return evalHashIndexExpression(left, index)
	default:
		return NewError(ErrorSymBolMap[CODE_PARSE_INDEX_OPERATOR_UNSUPPORTED](string(left.Type_Object())))
	}
}

func evalArrayIndexExpression(array, index Object) Object {
	arrObj := array.(*Array)
	idx := index.(*Integer).Value
	max := int64(len(arrObj.Elements) - 1)

	if idx < 0 || idx > max {
		return NilValue
	}

	return arrObj.Elements[idx]
}

func evalHashLiteral(node *HashLiteral, Env *Environment_of_environment) Object {
	pairs := make(map[HashKey]HashPair, len(node.Pairs))

	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, Env)
		if isError(key) {
			return key
		}

		hashKey, ok := key.(Hashable)
		if !ok {
			return NewError(ErrorSymBolMap[CODE_PARSE_HASHKEY_ERROR](string(key.Type_Object())))
		}

		value := Eval(valueNode, Env)
		if isError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = HashPair{
			Key:   key,
			Value: value,
		}
	}

	return &Hash{Pairs: pairs}
}

func evalHashIndexExpression(left, index Object) Object {
	key, ok := index.(Hashable)
	if !ok {
		return NewError(ErrorSymBolMap[CODE_PARSE_HASHKEY_ERROR](string(index.Type_Object())))
	}

	hashObj := left.(*Hash)
	if pair, exists := hashObj.Pairs[key.HashKey()]; exists {
		return pair.Value
	}
	return NilValue
}

func evalIdent(node *Ident, Env *Environment_of_environment) Object {
	if val, ok := Env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := Builtins[node.Value]; ok {
		return builtin
	}
	return NewError(ErrorSymBolMap[CODE_PARSE_IDENTIFIER_ERROR](node.Value))
}

func evalExpressions(exprs []Expression, Env *Environment_of_environment) []Object {
	result := make([]Object, 0, len(exprs))

	for _, expr := range exprs {
		evaluated := Eval(expr, Env)
		if isError(evaluated) {
			return []Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func EvalAssignModify(Assign *AssignmentStatement, Env *Environment_of_environment) Object {
	evaled := Eval(Assign.Value, Env)
	if isError(evaled) {
		return evaled
	}
	if Assign.Operator == "+=" || Assign.Operator == "-=" || Assign.Operator == "*=" || Assign.Operator == "/=" {
		currentvalue, ok := Env.Get(Assign.Name.Value)
		if !ok {
			return NewError("%s might not exist or is unknown to the parser -> \n SUGGESTION |+ allow %s = ...", Assign.Name.String(), Assign.Name.String())
		}
		res := evalInfixExpression(Assign.Operator, currentvalue, evaled)
		if isError(res) {
			fmt.Printf("Error during handle and evaluation to symbol %s %s ", Assign.Operator, res.Inspect())
			return res
		}
		Env.Set(Assign.Name.String(), res)
	} else if Assign.Operator == "=" {
		Env.Set(Assign.Name.String(), evaled)
	}
	return evaled
}

func evalProgram(program *Program, Env *Environment_of_environment) Object {
	var result Object

	for _, stmt := range program.Statements {
		result = Eval(stmt, Env)

		switch result := result.(type) {
		case *ReturnValue:
			return result.Value
		case *Error:
			return result
		}
	}

	return result
}

func nativeBoolToBooleanObject(input bool) *Boolean_Object {
	if input {
		return TrueValue
	}
	return FalseValue
}

func evalPrefixExpression(operator string, right Object) Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		// Eventually add support for prefix and infix seperate errors
		return NewError(ErrorSymBolMap[CODE_PARSE_OPERATOR_ERROR](operator))
	}
}

func evalBangOperatorExpression(right Object) Object {
	if right == NilValue || right == FalseValue {
		return TrueValue
	}
	return FalseValue
}

func evalMinusPrefixOperatorExpression(right Object) Object {
	switch right := right.(type) {
	case *Integer:
		return &Integer{Value: -right.Value}
	case *Float:
		return &Float{Value: -right.Value}
	default:
		return NewError("unknown operator: -%s", right.Type_Object())
	}
}

func EnvalBooleanInfixEXPRESSION_NODE(operator string, left, right Object) Object {
	lft := &String{Value: string(left.Inspect())}
	rft := &String{Value: string(right.Inspect())}
	switch operator {
	case "<":
		return evalStringInfixExpression(operator, lft, rft)
	case "<=":
		return evalStringInfixExpression(operator, left, rft)
	case ">":
		return evalStringInfixExpression(operator, lft, rft)
	case ">=":
		return evalStringInfixExpression(operator, left, rft)
	default:
		return NewError("Unknown operator: %s %s %s", lft.Type_Object(), operator, rft.Type_Object())

	}
}

func evalInfixExpression(operator string, left, right Object) Object {
	switch operator {
	}
	switch {
	case operator == "+" && left.Type_Object() == HashType && right.Type_Object() == HashType:
		lv := left.(*Hash).Pairs
		rv := right.(*Hash).Pairs
		pair := make(map[HashKey]HashPair)
		for idx, x := range lv {
			pair[idx] = x
		}
		for idx, c := range rv {
			pair[idx] = c
		}
		return &Hash{Pairs: pair}
	case operator == "+" && left.Type_Object() == ArrayType && right.Type_Object() == ArrayType:
		lv, rv := left.(*Array).Elements, right.(*Array).Elements
		elements := make([]Object, len(lv)+len(rv))
		elements = append(elements, rv...)
		return &Array{Elements: elements}
	case operator == "*" && left.Type_Object() == ArrayType && right.Type_Object() == IntegerType:
		lv, rv := left.(*Array).Elements, int(right.(*Integer).Value)
		elements := lv
		for idx := rv; idx > 1; idx-- {
			elements = append(elements, lv...)
		}
		return &Array{Elements: elements}
	case operator == "*" && left.Type_Object() == IntegerType && right.Type_Object() == ArrayType:
		lv, rv := int(left.(*Integer).Value), right.(*Array).Elements
		elements := rv
		for idx := lv; idx > 1; idx-- {
			elements = append(elements, rv...)
		}
		return &Array{Elements: elements}
	case operator == "*" && left.Type_Object() == STRING && right.Type_Object() == IntegerType:
		lv := left.(*String).Value
		rv := right.(*Integer).Value
		return &String{Value: strings.Repeat(lv, int(rv))}
	case operator == "*" && left.Type_Object() == IntegerType && right.Type_Object() == StringType:
		lv, rv := left.(*Integer).Value, right.(*String).Value
		return &String{Value: strings.Repeat(rv, int(lv))}
	case left.Type_Object() == IntegerType && right.Type_Object() == IntegerType:
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type_Object() == FloatType && right.Type_Object() == FloatType:
		return evalFloatInfixExpression(operator, left, right)
	case left.Type_Object() == FloatType && right.Type_Object() == IntegerType:
		return evalFloatIntegerInfixExpression(operator, left, right)
	case left.Type_Object() == IntegerType && right.Type_Object() == FloatType:
		return EvalSpecialFloatInfix(operator, left, right)
	case left.Type_Object() == StringType && right.Type_Object() == StringType:
		return evalStringInfixExpression(operator, left, right)
	case operator == "&&":
		return nativeBoolToBooleanObject(OBJECT_CONV_TO_NATIVE_BOOL(left) && OBJECT_CONV_TO_NATIVE_BOOL(right))
	case operator == "||":
		return nativeBoolToBooleanObject(OBJECT_CONV_TO_NATIVE_BOOL(left) || OBJECT_CONV_TO_NATIVE_BOOL(right))
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	case left.Type_Object() == BooleanType && right.Type_Object() == BooleanType:
		return EnvalBooleanInfixEXPRESSION_NODE(operator, left, right)
	case left.Type_Object() != right.Type_Object():
		return NewError(ErrorSymBolMap[CODE_PARSE_TYPE_ERROR](string(left.Type_Object()), string(right.Type_Object()), operator))
	default:
		return NewError(ErrorSymBolMap[CODE_PARSE_OPERATOR_ERROR](operator, string(left.Type_Object()), string(right.Type_Object())))
	}
}

// FLOATS ARE NOT INTEGERS THEY ARE THEIR OWN DATA TYPE | Modify later
func EvalSpecialFloatInfix(operator string, left, right Object) Object {
	leftVal := left.(*Float).Value
	rightVal := right.(*Float).Value
	switch operator {
	case "+":
		return &Float{Value: leftVal + rightVal}
	case "+=":
		return &Float{Value: leftVal + rightVal}
	case "-":
		return &Float{Value: leftVal - rightVal}
	case "-=":
		return &Float{Value: leftVal - rightVal}
	case "*":
		return &Float{Value: leftVal * rightVal}
	case "*=":
		return &Float{Value: leftVal * rightVal}
	case "**":
		return &Float{Value: math.Pow(leftVal, rightVal)}
	case "/":
		return &Float{Value: leftVal / rightVal}
	case "/=":
		return &Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return NewError("unknown operator: %s %s %s",
			left.Type_Object(), operator, right.Type_Object())
	}
}

func evalFloatIntegerInfixExpression(operator string, left, right Object) Object {
	leftVal := left.(*Float).Value
	rightVal := float64(right.(*Integer).Value)
	switch operator {
	case "+":
		return &Float{Value: leftVal + rightVal}
	case "+=":
		return &Float{Value: leftVal + rightVal}
	case "-":
		return &Float{Value: leftVal - rightVal}
	case "-=":
		return &Float{Value: leftVal - rightVal}
	case "*":
		return &Float{Value: leftVal * rightVal}
	case "*=":
		return &Float{Value: leftVal * rightVal}
	case "**":
		return &Float{Value: math.Pow(leftVal, rightVal)}
	case "/":
		return &Float{Value: leftVal / rightVal}
	case "/=":
		return &Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return NewError("unknown operator: %s %s %s",
			left.Type_Object(), operator, right.Type_Object())
	}
}

func evalIntegerInfixExpression(operator string, left, right Object) Object {
	leftVal := left.(*Integer).Value
	rightVal := right.(*Integer).Value

	switch operator {
	case "%":
		return &Integer{Value: leftVal % rightVal}
	case "+":
		return &Integer{Value: leftVal + rightVal}
	case "-":
		return &Integer{Value: leftVal - rightVal}
	case "-=":
		return &Integer{Value: leftVal - rightVal}
	case "*":
		return &Integer{Value: leftVal * rightVal}
	case "**":
		return &Integer{Value: int64(math.Pow(float64(leftVal), float64(rightVal)))}
	case "*=":
		return &Integer{Value: leftVal * rightVal}
	case "/":
		return &Integer{Value: leftVal / rightVal}
	case "/=":
		return &Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	case "..":
		len := int(rightVal-leftVal) + 1
		array := make([]Object, len)
		i := 0
		for i < len {
			array[i] = &Integer{Value: leftVal}
			leftVal++
			i++
		}
		return &Array{Elements: array}
	default:
		println("DEBUG -> OPERATOR FOR INFIX EXPRESSION -> ", operator)

		return NewError(ErrorSymBolMap[CODE_PARSE_OPERATOR_ERROR](operator, string(left.Type_Object()), string(right.Type_Object())))
	}
}

func evalFloatInfixExpression(operator string, left, right Object) Object {
	var leftVal, rightVal float64

	switch left := left.(type) {
	case *Integer:
		leftVal = float64(left.Value)
	case *Float:
		leftVal = left.Value
	default:
		return NewError(ErrorSymBolMap[CODE_PARSE_OPERATOR_ERROR](operator, string(left.Type_Object()), string(right.Type_Object())))
	}

	switch right := right.(type) {
	case *Integer:
		rightVal = float64(right.Value)
	case *Float:
		rightVal = right.Value
	default:
		return NewError(ErrorSymBolMap[CODE_PARSE_OPERATOR_ERROR](operator, string(left.Type_Object()), string(right.Type_Object())))
	}

	switch operator {
	case "+":
		return &Float{Value: leftVal + rightVal}
	case "+=":
		return &Float{Value: leftVal + rightVal}
	case "-":
		return &Float{Value: leftVal - rightVal}
	case "-=":
		return &Float{Value: leftVal - rightVal}
	case "*":
		return &Float{Value: leftVal * rightVal}
	case "**":
		return &Float{Value: math.Pow(leftVal, rightVal)}
	case "*=":
		return &Float{Value: leftVal * rightVal}
	case "/":
		return &Float{Value: leftVal / rightVal}
	case "/=":
		return &Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return NewError(ErrorSymBolMap[CODE_PARSE_OPERATOR_ERROR](operator, string(left.Type_Object()), string(right.Type_Object())))
	}
}

func evalStringInfixExpression(operator string, left, right Object) Object {
	leftVal := left.(*String).Value
	rightVal := right.(*String).Value

	switch operator {
	case "+":
		return &String{Value: leftVal + rightVal}
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case "+=":
		return &String{Value: leftVal + rightVal}
	default:
		return NewError(ErrorSymBolMap[CODE_PARSE_OPERATOR_ERROR](operator, string(left.Type_Object()), string(right.Type_Object())))
	}
}

func evalBlockStatement(block *BlockStatement, Env *Environment_of_environment) Object {
	var result Object

	for _, stmt := range block.Statements {
		result = Eval(stmt, Env)
		if result != nil {
			rt := result.Type_Object()
			if rt == ReturnValueType || rt == ErrorType {
				return result
			}
		}
	}

	return result
}

func evalIfExpression(ie *ConditionalExpression, Env *Environment_of_environment) Object {
	condition := Eval(ie.Condition, Env)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(ie.Consequence, Env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, Env)
	}
	return NilValue
}

func EvalWhileCondition(while *WhileCondition, Env *Environment_of_environment) Object {
	condition := Eval(while.Condition, Env)
	if isError(condition) {
		return condition
	}
	for isTruthy(condition) {
		Eval(while.Body, Env)
	}
	return &Nil{}
}
