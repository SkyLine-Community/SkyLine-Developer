package SkyLine_Backend

func (e *Environment_of_environment) SetConstants(name string, value Object) Object {
	e.Store[name] = value
	e.ROM[name] = true
	return value
}
