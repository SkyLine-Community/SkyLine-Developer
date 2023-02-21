package SkyLine_Backend

func NewEnclosedEnvironment(outer *Environment_of_environment) *Environment_of_environment {
	return &Environment_of_environment{
		Store: make(map[string]Object),
		Outer: outer,
	}
}
