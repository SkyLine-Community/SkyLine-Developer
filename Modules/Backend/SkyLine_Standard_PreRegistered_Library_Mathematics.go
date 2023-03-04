package SkyLine_Backend

import (
	"math"
	"strconv"
)

func mathAbs(args ...Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		v := arg.Value
		if v < 0 {
			v = v * -1
		}
		return &Integer{Value: v}
	case *Float:
		v := arg.Value
		if v < 0 {
			v = v * -1
		}
		return &Float{Value: v}
	default:
		return NewError("argument to `math.abs` not supported, got=%s",
			args[0].Type_Object())
	}
}

func mathCos(args ...Object) Object {
	if len(args) != 1 {
		return NewError("Wrong number of arguments. got=% but need=1", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Integer{Value: int64(math.Cos(float64(arg.Value)))}
	case *Float:
		return &Float{Value: math.Cos(arg.Value)}
	default:
		return NewError("Argument to `math.cos` must be integer or float")
	}
}

func mathSin(args ...Object) Object {
	if len(args) != 1 {
		return NewError("Wrong number of arguments. got=%d but need=1", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Integer{Value: int64(math.Sin(float64(arg.Value)))}
	case *Float:
		return &Float{Value: float64(math.Sin(arg.Value))}
	default:
		return NewError("Argument in call to `math.cos` must be integer or float")
	}
}

func mathTan(args ...Object) Object {
	if len(args) != 1 {
		return NewError("Wrong number of arguments, got=%d bnut need=1", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Integer{Value: int64(math.Tan(float64(arg.Value)))}
	case *Float:
		return &Float{Value: math.Tan(arg.Value)}
	default:
		return NewError("Argument in call to `math.tan` must be integer or float")
	}
}

func mathSqrt(args ...Object) Object {
	if len(args) != 1 {
		return NewError("Wrong number of arguments, got=%d but need=1", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Integer{Value: int64(math.Sqrt(float64(arg.Value)))}
	case *Float:
		return &Float{Value: math.Sqrt(arg.Value)}
	default:
		return NewError("Argument in call to `math.sqrt` must be integer or float")
	}
}

func Cbrt(args ...Object) Object {
	switch arg := args[0].(type) {
	case *Integer:
		f, _ := strconv.Atoi(arg.Inspect())
		z := f / 3.0
		for i := 0; i < 10; i++ {
			z = z - ((z*z*z - f) / (3 * z * z))
		}
		return &Integer{Value: int64(z)}
	case *Float:
		f, _ := strconv.Atoi(arg.Inspect())
		z := f / 3.0
		for i := 0; i < 10; i++ {
			z = z - ((z*z*z - f) / (3 * z * z))
		}
		return &Float{Value: float64(z)}
	default:
		return NewError("Argument in call to `math.cbrt` must be  integer or float")
	}
}

// speed is distance travled divided by time taken
