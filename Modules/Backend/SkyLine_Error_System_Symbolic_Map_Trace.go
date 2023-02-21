package SkyLine_Backend

import "fmt"

var ErrorSymBolMap = map[string]func(Arguments ...string) string{
	CODE_PARSE_FLOAT_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Could not parse (%s) as type float ", Arguments[0])
	}, // Error | Could not parse float value
	CODE_PARSE_INT_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Could not parse (%s) as type integer ", Arguments[0])
	}, // Error | Could not parse integer value
	CODE_PARSE_STRING_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Could not parse (%s) as type string", Arguments[0])
	}, // Error | Could not parse string value
	CODE_PARSE_BOOL_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Could not parse (%s) as type boolean", Arguments[0])
	}, // Error | Could not parse boolean value
	CODE_PARSE_NULL_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Could not parse (%s) as type null", Arguments[0])
	}, // Error | Could not parse NULL value
	CODE_PARSE_ARRAY_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Could not parse (%s) as type array", Arguments[0])
	}, // Error | Could not parse ARR
	CODE_PARSE_HASH_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Could not parse (%s) as type hash", Arguments[0])
	}, // Error | Could not parse HASH
	CODE_PARSE_HASHKEY_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Could not make  (%s) a useable hash key ", Arguments[0])
	}, // Error | Could not parse HashKey
	CODE_PARSE_TYPE_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Data Type mismatch in concentration, variable argument, call argument or function argument (Mismatch of type: %s and %s with operator (%s)) ", Arguments[0], Arguments[1], Arguments[2])
	}, // Er
	CODE_PARSE_OPERATOR_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Invalid operator (%s) in code block | %s %s %s | run Skyline__('OPERATORS') for more info", Arguments[0], Arguments[1], Arguments[0], Arguments[2])
	}, //
	CODE_PARSE_IDENTIFIER_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("No parse function found for identifier (%s) ", Arguments[0])
	}, //
	CODE_PARSE_FUNCTION_ARGUMENTS_NOT_ENOUGH_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Function (%s) Does not have enough arguments in call to function or method Arguments -> Given(%s), Requires(%s)", Arguments[0], Arguments[1], Arguments[2])
	}, //
	CODE_PARSE_MACRO_INVALID_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Macro (%s) may not exist or may not be currently configured with the modifier", Arguments[0])
	}, //
	CODE_PARSE_INDEX_OPERATOR_UNSUPPORTED_WITHIN_KEY_NOTE_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Index operator used to index the array or hash is not currently supported with token literal (%s) ", Arguments[0])
	}, //
	CODE_PARSE_AST_MODIFICATION_TO_MACRO_UNSUPPORTED_METHOD_ERROR: func(Arguments ...string) string {
		return fmt.Sprintf("Modification unsupported in call to MacroExpansion, currently only returning AST (Abstract Syntax Tree) nodes are supported from macros in (%s) ", Arguments[0])
	}, //
	CODE_NO_FUNCTIONS_OR_SYMBOLS_LOADED: func(Arguments ...string) string {
		return fmt.Sprintf("SkyLine Debug: The file loaded must not have called any symbols, methods etc as the code was not returning any data FNAME(%s) ", Arguments[0])
	}, //
	CODE_PARSE_INDEX_OPERATOR_UNSUPPORTED: func(Arguments ...string) string {
		return fmt.Sprintf("Invalid index operator which is unsupported (%s) ", Arguments[0])
	},
	CODE_WRONG_NUMBER_OF_ARGUMENTS: func(Arguments ...string) string {
		return fmt.Sprintf("Wrong number of arguments for builtin function (%s) which requires %s argument(s) but you gave %s argument(s)",
			Arguments[0],
			Arguments[1],
			Arguments[2],
		)
	}, ////////////////////
	// File integrity checks
	CODE_PREFIX_PARSE_FUNCTION_INVALID_OR_UNFOUND_WITHIN_PARSER_AND_INTERPRETRR: func(Arguments ...string) string {
		return fmt.Sprintf("Function or Method by name (%s) undefined", Arguments[0])
	}, //
	CODE_EXPECT_PEEK_ERROR_DURING_CALL_TO_PEEK: func(Arguments ...string) string {
		return fmt.Sprintf("Token Expection Error: Unexpected token (%s) expecting (%s)", Arguments[0], Arguments[1])
	}, //
	CODE_FILE_INTEGRITY_FILE_INVALID_MUST_BE_CSC_FILE: func(Arguments ...string) string {
		return fmt.Sprintf("SkyLine File Integrity: SkyLine can not process file (%s) because it must end in .csc, please ensure the filename is correct", Arguments[0])
	}, //
	CODE_FILE_INTEGRITY_FILE_INVALID_MUST_NOT_BE_DIRECTORY_DIR_UNSUPPORTED: func(Arguments ...string) string {
		return fmt.Sprintf("SkyLine File Integrity: SkyLine can not process (%s) because it is not a file, rather it is a directory, please ensure during import, require, include or carry these are real .csc files and not directories", Arguments[0])
	}, //
	CODE_FILE_INTEGRITY_FILE_INVALID_MUST_HAVE_CODE_OR_LOGIC_INSIDE_FILE_NULL: func(Arguments ...string) string {
		return fmt.Sprintf("SkyLine File Integrity: SkyLine refused to run file ( %s ) through the parser because the file does not have any code SEC WARNING...", Arguments[0])
	}, //
	CODE_FILE_INTEGRITY_FILE_INVALID_FAILED_TO_IMPORT_OR_OPEN_FILE: func(Arguments ...string) string {
		return fmt.Sprintf("SkyLine Import Integrity: Failed to import, require, include or carry %s, this file for some reason did not want to be oppened: SEC WARNING....", Arguments[0])
	}, //
	CODE_FILE_INTEGRITY_FILE_INVALID_FAILED_TO_STAT: func(Arguments ...string) string {
		return fmt.Sprintf("SkyLine File Integrity: Failed to stat (%s) for some reason when calling FileCurrent.New() the stat loader for file integrity has failed", Arguments[0])
	}, //
	CODE_FILE_INTEGRITY_FILE_INVALID_FILE_NAME_WAS_EMPTY_OR_NULL_CHEC_INPUT: func(Arguments ...string) string {
		return fmt.Sprintf("SkyLine New File Integrity: Failed to load file, file for input was empty, this is a weird error, how did you even get here? %s", Arguments[0])
	}, //
	CODE_FILE_INTEGRITY_FILE_INVALID_FILE_DOES_NOT_EXIST: func(Arguments ...string) string {
		return fmt.Sprintf("SkyLine File Integrity: File (%s) Is not an existing file, please check and verify your file names before you continue to run code", Arguments[0])
	}, //
	CODE_FILE_FAILED_INJECTION_FILE_FAILED_TO_LINK_DUE_TO_NULLERR: func(Arguments ...string) string {
		return fmt.Sprintf("SkyLine Linker: Linker attempted to inject the imported code from (%s) into the current runtime file and failed, the file that you tried to import is empty, please ensure data is in the file before trying to run it through the SkyLine interpretr", Arguments[0])
	}, //
	CODE_FILE_MUST_HAVE_NEW_FUNCTION_AND_METHOD_CALLED_DEVELOPER_ERROR_IN_SYMBOL: func(Arguments ...string) string {
		return fmt.Sprintf("SkyLine Dev Parser ERROR: SkyLine could not find the parser function for the given symbol ( %s ) ", Arguments[0])
	}, //
	CODE_FILE_FAILED_USING_INPUT_OUTPUT_READER_AND_UTILITY_FILE_ISSUE: func(Arguments ...string) string {
		return fmt.Sprintf("SkyLine File Integrity: File (%s) Failed to load from IO due to IO error (%s) ", Arguments[0], Arguments[1])
	}, // IOUTIL error
	CODE_SERVER_FAILED_TO_RESPOND: func(Arguments ...string) string {
		return "SkyLine Server: Failed to connect or grab data from the server, The server may have been disconnected Or may have failed to connect to the required ports"
	}, // SERVER FAIL

}

// Color list by OS

// ______LINUX OPERATING SYSTEMS__________

const (
	ERROR_RED = "\033[38;5;160m"
	ERROR_MSG = "\033[38;5;123m"
)
