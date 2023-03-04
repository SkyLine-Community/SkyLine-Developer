package SkyLine_Backend

func NewEnvironment() *Environment_of_environment {
	return &Environment_of_environment{
		Store: make(map[string]Object),
		Outer: nil,
		ROM:   make(map[string]bool),
	}
}
