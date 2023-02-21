package SkyLine_Backend

var Builtins = map[string]*Builtin{}

func RegisterBuiltin(name string, fun BuiltinFunction) {
	Builtins[name] = &Builtin{Fn: fun}
}
