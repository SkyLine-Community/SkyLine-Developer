package SkyLine_Backend

func (e *Environment_of_environment) Get(name string) (Object, bool) {
	obj, exists := e.Store[name]
	if !exists && e.Outer != nil {
		obj, exists = e.Outer.Get(name)
	}
	return obj, exists
}
