package SkyLine_Backend

import (
	"fmt"
	"log"
	SkyLine_External_Forensics "main/Modules/StandardLibraryExternal/Forensics"
	"os"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//
// This is the registry module, this module apart of SkyLine_Backend allows us to register standard library based functions which are called with
//
// class.functionname
//
// this is pretty simple to understand however class and module keywords have not been implemented which means that you can not register your
// own custom standard module / library. Asides from that, we use the init() function because init functions will always run or be called before
// the main() function in go. Using registers under the init function we can ensure the environment has standard functions registered and placed
// into the environment before it is fully started and the input program is parsed. This eliminates the need to import("math") however in the
// further future import keywords will need to be added for standard library functions. This is becausethe bigger our standard library gets the
// more imports will need to be added and the harder the program will be to parse. Currently, due to the factor of how small the standard library
// is, it is not that bad to register the built in functions before a new environment for the input program is started which means it will not slow
// down runtime. However, as this again gets bigger we will need to eliminate registering before runtime unless they are standard functions such as
// .str, .int, integer, boolean, empty?, nil?, carries?, exported? etc which allow for a much heavier use case and do not require imports
// Using the import keyword will give the user the option to allow the program to import and register the standard library functions before
// runtime and parsing. This may cause collisions within the environment so we can actually cause another keyword to exist known as "register"
// followed by the library name. This keyword may be called like so register("math") pr register<<"math">> for a much more complex and parsed
// syntax. Allowing both register and import keywords allow the user to register the library functions before runtime and import files before
// runtime.
//
//
// - Mon 27 Feb 2023 10:23:16 PM EST
//
// as of the given date and time of writing this, SkyLine now will ask you to register the library before you use them if it is standard
// this includes crypt, math, net, http and much other built in libraries used within the SkyLine programming language
//
//
// Lib type: This is the external library registration which means any library that is developed in PATH where
//
// PATH = SkyLine/Modules/StandardLibraryExternal
//
// is a external library that does not rely on the functions and types within this current filepath for modules. Rather than importing them
// we can register them as standard libraries

////////////////// FORENSICS LIBRARY AND AUTOMATION

var metadata SkyLine_External_Forensics.PNG_Meta

func RunMeta(args ...Object) Object {
	dat, err := os.Open(args[0].Inspect())
	if err != nil {
		log.Fatal(err)
	}
	defer dat.Close()
	bReader, err := SkyLine_External_Forensics.Process_Given_Image(dat)
	if err != nil {
		log.Fatal(err)
	}
	metadata.Metadata(bReader)
	return &Nil{}
}

var SessionSets SkyLine_External_Forensics.Settings

func RunSettings(args ...Object) Object {
	key := args[0].Inspect()
	out := args[1].Inspect()
	in := args[2].Inspect()
	filemode := 0
	offset := args[3].Inspect()
	payload := args[4].Inspect()
	chunktoinject := args[5].Inspect()
	SessionSets.Settings_Inject_New(
		key,
		out,
		in,
		fmt.Sprint(filemode),
		"false",
		"false",
		offset,
		payload,
		chunktoinject,
	)
	return &Nil{}
}

func RunInject(args ...Object) Object {
	dat, err := os.Open(args[0].Inspect())
	if err != nil {
		log.Fatal(err)
	}
	defer dat.Close()
	bReader, err := SkyLine_External_Forensics.Process_Given_Image(dat)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(SessionSets.ImageOffset)
	SessionSets.Injection_Standard_Payload(bReader)
	return &Boolean_Object{Value: true}
}

func RegisterForensics() {
	RegisterBuiltin("forensics.meta", func(env *Environment_of_environment, args ...Object) Object {
		return (RunMeta(args...))
	})
	RegisterBuiltin("forensics.PngSettingsNew", func(env *Environment_of_environment, args ...Object) Object {
		return (RunSettings(args...))
	})
	RegisterBuiltin("forensics.InjectPNG", func(env *Environment_of_environment, args ...Object) Object {
		return (RunInject(args...))
	})
}
