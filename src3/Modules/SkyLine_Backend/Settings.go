package SkyLine

import (
	"flag"
	"fmt"
	"runtime"
)

type UserInterpretData struct {
	OperatingSystem             string
	OperatingSystemArchitecture string
}

var (
	Help             = flag.Bool("help", false, "Load help module")
	ErrorsTrace      = flag.Bool("trace", false, "Load tracer module for errors, or if script empty output panic or recovery")
	SourceFile       = flag.String("source", "", "Load source code file into SkyLine")
	Bnn              = flag.Bool("bout", false, "If true will output the SkyLine banner when running a code file")
	Server           = flag.Bool("server", false, "If true will load the SkyLine local server")
	CompileWithGo    = flag.Bool("compile-with-go", false, "Compile with the interpreter but rather take the input of a source code file and compile it with the embedded interpreter")
	SettingsFileYAML = flag.String("settings-fileYAML", "", "Will load settings for SkyLine's interpreter from a YAML file")
	SettingsFileJSON = flag.String("settings-fileJSON", "", "Will load settings for SkyLine's interpreter from a JSON file")
)

func Banner() {
	fmt.Println("\x1b[H\x1b[2J\x1b[3J")
	U.OperatingSystem = runtime.GOOS
	U.OperatingSystemArchitecture = runtime.GOARCH
	switch runtime.GOOS {
	case "linux":
		fmt.Println("\t\t\t\t\t\t\t\t\t	 \033[38;5;51m┏━┓\x1b[0m")
		fmt.Println("\t\t\t\t\t\t\t\t\t	\033[38;5;56m┃\033[38;5;51m┃ ┃\x1b[0m")
		fmt.Println("\t\t\t\t\t\t\t\t\t    \033[38;5;56m━━━━┛\x1b[0m")
		fmt.Println("\t\t\t\t\t\t\t\t\t	\033[38;5;249mSky Line Interpreter| V 0.0.2")
		fmt.Print("\n\n\033[39m")
	default:
		fmt.Println("\t\t\t\t	 \u001b[38;5;51m┏━┓\u001b[0m")
		fmt.Println("\t\t\t\t	\u001b[38;5;56m┃\u001b[38;5;51m┃ ┃\u001b[0m")
		fmt.Println("\t\t\t\t    \u001b[38;5;56m━━━━┛\u001b[0m")
		fmt.Println("\t\t\t\t	\u001b[38;5;249mSky Line Interpreter| V 0.0.2")

	}
}
