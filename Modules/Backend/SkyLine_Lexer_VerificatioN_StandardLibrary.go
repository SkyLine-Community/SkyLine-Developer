package SkyLine_Backend

var ConstantIdents = map[string]bool{
	"math.abs":                 true,
	"math.cos":                 true,
	"math.sin":                 true,
	"math.sqrt":                true,
	"math.tan":                 true,
	"math.cbrt":                true,
	"math.rand":                true,
	"math.out":                 true,
	"crypt.hash":               true,
	"io.input":                 true,
	"io.clear":                 true,
	"io.box":                   true,
	"io.listen":                true,
	"forensics.new":            true,
	"forensics.meta":           true,
	"forensics.PngSettingsNew": true,
	"forensics.InjectPNG":      true,
}

var StandardLibNames = map[string]bool{
	"math":      true,
	"io":        true,
	"forensics": true,
	"os":        true,
	"http":      true,
	"crypt":     true,
}

var RegisterStandard = map[string]func(){
	"io":        RegisterIO,
	"math":      RegisterMath,
	"crypt":     RegisterCrypt,
	"forensics": RegisterForensics,
}

var Datatypes = []string{
	"string.",
	"float.",
	"object.",
	"hash.",
	"array.",
}
