package SkyLine_Backend

import "fmt"

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////// | The offical SkyLine error system, this file includes code systems, notes as big as this and messages as well as loggers and other statements which allow you to make a
///////  | much more verbose and advanced error system. This system uses a custom code and output system to define its errors. Do note that this is still in beta and needs to be
///////  | tested much further before deployment and that it is important to note that not all codes need to be used.
//////   |
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// Date  : Mon 27 Feb 2023 10:28:49 PM EST
// Codes :
// Types : Files, Parser errors, argument error, type errors, unknown operators, operational errors, mathematical errors, byte errors, binary errors, flow errors, systematic errors, os errors, support errors, network errors and more.
//
//
// The code system for SkyLine will differ based on the type of code and the type of system you are trying to troubleshoot, for example
// certain modules, functions, nodes and starters have reserved sets of codes which all mean different things and which have their own error messages.
// With every error based on the error and type of code if it is supported or required will allow for recomendations such as missing imports, semicolons etc
// below you will find a list of all the error codes and their reserved spots
//
// Startup        | -> 0-100
// Parser         | -> 100-300
// Evaluator      | -> 300-500
// Operations     | -> 500-700
// Environment    | -> 800-1000
// Function Calls | -> 1000-2000
// Network calls  | -> 2000-3000
// Mathematics    | -> 4000-5000
// Data types     | -> 5000-6000
// Server errors  | -> 6000-6500
// Arguments      | -> 6500-7000
//
//
// Each message and each code depending on the type of code and the catgeory increments completely differently and is all formatted differently. For example below you will find a
// list of the error code reservations and their incremented positions
//
//
// Startup        | -> 0-100         | Increments by 5
// Parser         | -> 100-300       | Increments by 5
// Evaluator      | -> 300-500       | Increments by 10
// Operations     | -> 500-700       | Increments by 20
// Environment    | -> 800-1000      | Increments by 30
// STD library    | -> 1000-2000     | Increments by 10
// Network calls  | -> 2000-3000     | Increments by 20
// Mathematics    | -> 4000-5000     | Increments by 80
// Data types     | -> 5000-6000     | Increments by 90
// Server errors  | -> 6000-10,000   | Increments by 100
// Arguments      | -> 10,000-10,500 | Increments by 100s
//
// These arguments all have their own respected values and all have their own unique values and codes as well
// for example if you were to make a HTTP POST request to a given domain and the endpoint failes a code within the
// range of 2000-3000 will be displayed because this counts as a network error within the system. Anything past 10,500
// is just wild and much more verbose and a complicated error or traceback. Codes should not ever reach past that point
// as that typically means there is a panic within the program, a fatal log, a dev issue on the backend or some weird issue.
// It is important to note that within each system there is no line or column based traceback system right now as this is
// a main priority and I would rather much prefer to focus on the functionality of the error system first before the backend
//
//
//
//
//
//
//
//
//
// Startup errors
const (
	ERROR_SOURCE_FLAG_DEFINED_FILE_DOES_NOT_EXIST = 5  // File does not exist
	ERROR_SOURCE_FLAG_DEFINED_FILE_IS_A_DIRECTORY = 10 // File is a directory
	ERROR_SOURCE_FLAG_DEFINED_FILE_IS_EMPTY       = 15 // File is empty, this is bad to run in some cases
	ERROR_SOURCE_FLAG_DEFINED_FILE_IS_INVALID_SRC = 20 // File is a invalid source code file ( does not end in .csc, .skyline, .sl, .srcsl, .datasl)
)

// Parser Errors
const (
	ERROR_DURING_PEEK_IN_PARSER              = 100 // Peek error, missing token after so and so character
	ERROR_FILE_INTEGRITY_BAD_FILE            = 105 // File from import or require was not a CSC/SkyLine source file
	ERROR_FILE_INTEGRITY_IS_DIRECTORY        = 110 // File from import or require was a directory
	ERROR_FILE_INTEGRITY_IS_EMPTY            = 115 // File from import or require was empty
	ERROR_FILE_INTEGRITY_DOES_NOT_EXIST      = 120 // File from import or require does not exist or was not found
	ERROR_FILE_INPUT_OUTPUT_BUFFER_FAILED    = 125 // File from import or require was not able to be read by the reader
	ERROR_TYPE_INTEGRITY_PARSE_INTEGER_ERROR = 130 // Parser failed to parse an integer of the supported types
)

// Evaluator Errors
const (
	ERROR_NOT_ENOUGH_ARGUMENTS_IN_CALL_TO_FUNCTION             = 300 // Arguments in call to function were not correct, function requires ... but unexpected to get ...
	ERROR_INVALID_OPERATOR_DURING_EVALUATION                   = 310 // Invalid prefix operator when supported includes (! and -)
	ERROR_INVALID_FUNCTION_NOT_FOUND_OR_UNSUPPORTED_TYPE       = 320 // Function or prefix parse function did not exist nor did its respective type exist
	ERROR_INVALID_INDEX_EXPRESSION_OR_UNSUPPORTED_INDEX        = 330 // Index operator unsupported or non existent
	ERROR_INVALID_HASH_KEY_COULD_NOT_BE_USED_AS_KEY            = 340 // Key within hash could not be hashable in hash literal
	ERROR_INVALID_HASH_KEY_COULD_NOT_PARSE_HASHKEY             = 350 // Key within hash in evaluation to hash literal could not be parsed
	ERROR_INVALID_IDENTIFIER_IDENTIFIER_WAS_NOT_FOUND_OR_KNOWN = 360 // Identifier was not found or located within the current environment
	ERROR_INVALID_DATATYPE_INIFX_OPERATION_LR                  = 370 // the type of the left value does not equal the type of the right value
)

// Environment Errors
const (
	ERROR_EXTEND_FUNCTION_NOT_ENOUGH_ARGUMENTS = 800 // Not enough arguments in call to function
	ERROR_CONSTANT_VALUE_ATTEMPT_TO_CHANGE     = 830 // Cannot change the value of constant x
)

// Standard Library Errors
const (
	ERROR_STANDARD_FUNCTION_DOES_NOT_HAVE_PROPER_ARGUMENTS         = 1000 // In call to standard library function ... does not have enough arguments, was given ... but requires ...
	ERROR_STANDARD_FUNCITON_REQUIRES_SO_AND_SO_TYPE_NOT_PROVIDED   = 1010 // In call to standard library function ... the given argument at (n) was provided to be the wrong data type
	ERROR_STANDARD_PRESET_CODE_ARGUMENT_WAS_NOT_FOUND_OR_SUPPORTED = 1020 // In call to standard library function ... the given argument which is a default argument was not found or expected
)

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////// | This next section here will define the maps and function calls necessary to operate with the error codes. These variables are maps
///////  | which result in a function that returns a new message and error setting called by and called from sprintf. These messages can be
///////  | of Object error type and will need to return a new error message and setting or recomendation.
///////  |
///////  | - Date: Tue 28 Feb 2023 04:37:17 PM EST
///////  |
///////  |
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// C_1 -> Map_Eval   : Eval        error specific map
// C_2 -> Map_Parser : Parser      error specific map
// C_3 -> Map_Std    : Stdlib      error specific map
// C_4 -> Map_Envi   : Environment error specific map
//

// type

type ErrorSystemReturn struct {
	Message    string // Message
	Code       int    // Error code
	Suggestion string // Suggested fix
}

func GenerateErrorDeveloper() ErrorSystemReturn {
	return ErrorSystemReturn{
		Message:    "Error system function call was not found, developer error ( not enough arguments in call to error map) ",
		Code:       ERROR_STANDARD_FUNCTION_DOES_NOT_HAVE_PROPER_ARGUMENTS,
		Suggestion: "Contact developers if this issue appears",
	}
}

// C4 | Environment errors during token identification
var Map_Envi = map[int]func(Arguments ...string) ErrorSystemReturn{
	ERROR_CONSTANT_VALUE_ATTEMPT_TO_CHANGE: func(Arguments ...string) ErrorSystemReturn {
		return ErrorSystemReturn{
			Message:    "Attempting to change constant value which is ILLEGAL",
			Code:       ERROR_CONSTANT_VALUE_ATTEMPT_TO_CHANGE,
			Suggestion: "Do not change constants",
		}
	},
	ERROR_EXTEND_FUNCTION_NOT_ENOUGH_ARGUMENTS: func(Arguments ...string) ErrorSystemReturn {
		return ErrorSystemReturn{
			Message:    "A certain function was called but was not given the proper argument counts.",
			Code:       ERROR_EXTEND_FUNCTION_NOT_ENOUGH_ARGUMENTS,
			Suggestion: "Meet the function argument requirements",
		}
	},
}

// C3 | Standard Library errors
var Map_STDLIB = map[int]func(Arguments ...string) ErrorSystemReturn{
	ERROR_STANDARD_FUNCTION_DOES_NOT_HAVE_PROPER_ARGUMENTS: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 3 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf("Standard function call to (%s) requires at least %s argument(s) but you have (%s) argument(s)",
					Arguments[0],
					Arguments[1],
					Arguments[2],
				),
				Code: ERROR_STANDARD_FUNCTION_DOES_NOT_HAVE_PROPER_ARGUMENTS,
				Suggestion: fmt.Sprintf("Suggestion : When calling function (%s) use only %s argument(s) to prevent this error", Arguments[0],
					Arguments[1]),
			}
		} else {
			return GenerateErrorDeveloper()
		}
	},
	ERROR_STANDARD_FUNCITON_REQUIRES_SO_AND_SO_TYPE_NOT_PROVIDED: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 4 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf("Standard library call to function (%s) was given the wrong data type at POS (%s) in the argument list, the data type of the variable given was %s but the function expects you to at least have %s",
					Arguments[0],
					Arguments[1],
					Arguments[2],
					Arguments[3],
				),
				Code:       ERROR_STANDARD_FUNCITON_REQUIRES_SO_AND_SO_TYPE_NOT_PROVIDED,
				Suggestion: "Use the correct data type the function requires",
			}
		} else {
			return GenerateErrorDeveloper()
		}
	},
	ERROR_STANDARD_PRESET_CODE_ARGUMENT_WAS_NOT_FOUND_OR_SUPPORTED: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 3 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"In call to standard library function %s, this function has a param that is pre set on the backend or static which requires a specific title. You gave this function the following argument %s in POS %s but this function requires a different argument which may not be supported",
					Arguments[0],
					Arguments[1],
					Arguments[2],
				),
				Code:       ERROR_STANDARD_PRESET_CODE_ARGUMENT_WAS_NOT_FOUND_OR_SUPPORTED,
				Suggestion: "Use the std.lookup function to lookup standard library functions and fill out the required details to lookup the functions details and requirements or suggestions",
			}
		} else {
			return GenerateErrorDeveloper()
		}
	},
}

// C2 | Parser error codes, this map will return error codes for the parser itself.
var Map_Parser = map[int]func(Arguments ...string) ErrorSystemReturn{
	ERROR_DURING_PEEK_IN_PARSER: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 2 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"When peeking or looking for the next token, the parser could not find the required token. Expected next token to be %s but rather found %s",
					Arguments[0],
					Arguments[1],
				),
				Code:       ERROR_DURING_PEEK_IN_PARSER,
				Suggestion: "Check to make sure your token is correct",
			}
		} else {
			return GenerateErrorDeveloper()
		}
	},
	ERROR_FILE_INTEGRITY_BAD_FILE: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"When parsing the file PARSE(%s), the file fails to become parsed or called due to the file being bad. ",
					Arguments[0],
				),
				Code:       ERROR_FILE_INTEGRITY_BAD_FILE,
				Suggestion: "File extension INVALID -> File types are .csc, .sl, .skyline, .SkyLine, .core, .seccore",
			}
		} else {
			return GenerateErrorDeveloper()
		}
	},
	ERROR_FILE_INTEGRITY_DOES_NOT_EXIST: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"When parsing for file (%s) the standard checks detected that this file did not exist within the given directory",
					Arguments[0],
				),
				Code:       ERROR_FILE_INTEGRITY_DOES_NOT_EXIST,
				Suggestion: fmt.Sprintf("Make sure file %s is correct ", Arguments[0]),
			}
		} else {
			return GenerateErrorDeveloper()
		}
	},
	ERROR_FILE_INTEGRITY_IS_DIRECTORY: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"When parsing for file (%s) the standard checks detected that this was not a file but rather a directory. Execution of CSC files within a directory is not supported in the current version...",
					Arguments[0],
				),
				Code:       ERROR_FILE_INTEGRITY_IS_DIRECTORY,
				Suggestion: "Provide a source code file, not a directory",
			}
		} else {
			return GenerateErrorDeveloper()
		}
	},
	ERROR_FILE_INTEGRITY_IS_EMPTY: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"When parsing for file (%s) the standard checks detected that this was a empty file, this is a security issue and the interpreter will not accept any file that is empty.",
					Arguments[0],
				),
				Code:       ERROR_FILE_INTEGRITY_IS_EMPTY,
				Suggestion: "Provide a source code file that has something inside of it and is not NULL",
			}
		} else {
			return GenerateErrorDeveloper()
		}
	},
	ERROR_FILE_INPUT_OUTPUT_BUFFER_FAILED: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"When parsing the file (%s) the Input and Output buffer when reading the file has failed for some reason as it gave this error (%s)",
					Arguments[0],
					Arguments[0],
				),
				Code:       ERROR_FILE_INPUT_OUTPUT_BUFFER_FAILED,
				Suggestion: "Provide a source code file that may not be corrupted, damaged, injected or infected with data or is being used by another process",
			}
		} else {
			return GenerateErrorDeveloper()
		}

	},
	ERROR_TYPE_INTEGRITY_PARSE_INTEGER_ERROR: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"When parsing integer (%s) the parser came across an error that made it bad, unsafe or made the parser error out when parsing the provided integer",
					Arguments[0],
				),
				Code:       ERROR_TYPE_INTEGRITY_PARSE_INTEGER_ERROR,
				Suggestion: "Provide a number variable that does not overflow the Integer64 data type indicating that it can be a real integer and can be parsed by SkyLine",
			}
		} else {
			return GenerateErrorDeveloper()
		}
	},
}

// C1 | Evaluation step errors

var Map_Eval = map[int]func(Arguments ...string) ErrorSystemReturn{
	ERROR_NOT_ENOUGH_ARGUMENTS_IN_CALL_TO_FUNCTION: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 3 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"In call to function %s, there were not enough arguments that were needed as this function requires %s argumen(s) but you gave it %s argument(s)",
					Arguments[0],
					Arguments[1],
					Arguments[2],
				),
				Code:       ERROR_NOT_ENOUGH_ARGUMENTS_IN_CALL_TO_FUNCTION,
				Suggestion: "Use the correct number of arguments",
			}
		} else {
			return GenerateErrorDeveloper()

		}
	},
	ERROR_INVALID_OPERATOR_DURING_EVALUATION: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Code:       ERROR_INVALID_OPERATOR_DURING_EVALUATION,
				Suggestion: "Use the supported operators",
				Message: fmt.Sprintf(
					"When evalutating a given operator (%s) the evaluator found that it was not supported, please check the laws of the evaluator",
					Arguments[0],
				),
			}
		} else {
			return GenerateErrorDeveloper()
		}
	},
	ERROR_INVALID_FUNCTION_NOT_FOUND_OR_UNSUPPORTED_TYPE: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Code:       ERROR_INVALID_FUNCTION_NOT_FOUND_OR_UNSUPPORTED_TYPE,
				Suggestion: "Make sure evaluator has this valid",
				Message: fmt.Sprintf(
					"Function type (%s) was not found, this is bad please make sure the function is the supported type or this feature was implemented",
					Arguments[0],
				),
			}

		} else {
			return GenerateErrorDeveloper()

		}
	},
	ERROR_INVALID_INDEX_EXPRESSION_OR_UNSUPPORTED_INDEX: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Code: ERROR_INVALID_INDEX_EXPRESSION_OR_UNSUPPORTED_INDEX,
				Message: fmt.Sprintf(
					"Index operator (%s) unsupported or non existent",
					Arguments[0],
				),
				Suggestion: "Make sure the index operator is supported ",
			}

		} else {
			return GenerateErrorDeveloper()

		}
	},
	ERROR_INVALID_HASH_KEY_COULD_NOT_BE_USED_AS_KEY: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Code:       ERROR_INVALID_HASH_KEY_COULD_NOT_BE_USED_AS_KEY,
				Message:    fmt.Sprintf("Key (%s) was not able to be used as a valid hash key", Arguments[0]),
				Suggestion: "make sure the hash key you gave is correct and stable",
			}
		} else {
			return GenerateErrorDeveloper()
		}
	},
	ERROR_INVALID_HASH_KEY_COULD_NOT_PARSE_HASHKEY: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Code:       ERROR_INVALID_HASH_KEY_COULD_NOT_PARSE_HASHKEY,
				Message:    fmt.Sprintf("Could not parse hash key (%s)", Arguments[0]),
				Suggestion: "Make sure this key can be used within the hash",
			}
		} else {
			return GenerateErrorDeveloper()
		}
	},
	ERROR_INVALID_IDENTIFIER_IDENTIFIER_WAS_NOT_FOUND_OR_KNOWN: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 2 {
			return ErrorSystemReturn{
				Code:       ERROR_INVALID_IDENTIFIER_IDENTIFIER_WAS_NOT_FOUND_OR_KNOWN,
				Suggestion: "The identifier you tried to parse was not assigned or registered to the environment, check if it exists",
				Message:    fmt.Sprintf("The identifier (%s) of type (%s) could not be found or was not known to the interpreters environment", Arguments[0], Arguments[1]),
			}

		} else {
			return GenerateErrorDeveloper()
		}
	},
	ERROR_INVALID_DATATYPE_INIFX_OPERATION_LR: func(Arguments ...string) ErrorSystemReturn {
		// needs left data type and right data type
		if len(Arguments) == 2 {
			return ErrorSystemReturn{
				Code:       ERROR_INVALID_IDENTIFIER_IDENTIFIER_WAS_NOT_FOUND_OR_KNOWN,
				Suggestion: "Make sure the type of the left value is equal to the type of the right value",
				Message: fmt.Sprintf(
					"During evaluation of the infix operation, the data type of the left value which yeilds (%s) must be equal to the type of the right value (%s) which in this case is not true, fix it.",
					Arguments[0],
					Arguments[1],
				),
			}

		} else {
			return GenerateErrorDeveloper()
		}
	},
}
