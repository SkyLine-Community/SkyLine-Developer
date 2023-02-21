package SkyLine_Backend

import "fmt"

func isTruthy(obj Object) bool {
	return obj != NilValue && obj != FalseValue
}

func NewError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj Object) bool {
	return obj != nil && obj.Type_Object() == ErrorType
}
