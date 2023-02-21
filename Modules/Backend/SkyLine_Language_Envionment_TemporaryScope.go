package SkyLine_Backend

func NewTempScop(outer *Environment_of_environment, keys []string) *Environment_of_environment {
	env := NewEnvironment()
	env.Outer = outer
	env.permit = keys
	return env
}
