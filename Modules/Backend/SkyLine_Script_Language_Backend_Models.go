package SkyLine_Backend

import (
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

///////////////////////////////////////////////////////////////////////////
//// 						Tokenizer types                          ////

type Token_Type string

type Token struct {
	Token_Type Token_Type
	Literal    string
}

const (
	IMPORT          = "IMPORT"
	REGISTER        = "REGISTER "
	FOREACH         = "FOREACH"
	IN              = "IN"       // Foreach .. in .... |
	ILLEGAL         = "ILLEGAL"  // Illegal character  | Implemented
	EOF             = "EOF"      // End Of File        | Implemented
	IDENT           = "IDENT"    // Identifier         | Implimented
	INT             = "INT"      // TYPE integer       | Implemented
	FLOAT           = "FLOAT"    // TYPE float         | Implemented
	STRING          = "STRING"   // TYPE string        | Implemented
	UNSIGNED        = "UNSIGNED" // TYPE unsigned int  | Implemented
	CONSTANT        = "CONST"    // Constant           | Implemented
	FUNCTION        = "FUNCTION" // Function           | Implemented
	LET             = "LET"      // let statement      | Implemented
	TRUE            = "TRUE"     // boolean type true  | Implemented
	FALSE           = "FALSE"    // boolean type false | Implemented
	IF              = "IF"       // If statement       | Implemented
	ELSE            = "ELSE"     // Else statement     | Implemented
	RETURN          = "RETURN"   // return statement   | Implemented
	SWITCH          = "SWITCH"   // Switch statement   | Implemented
	CASE            = "CASE"     // Case statement 	   | Implemented
	WHILE           = "WHILE"    // Declare while loop | Implemented
	LTEQ            = "<="       // LT or equal to     | Implemented
	GTEQ            = ">="       // GT or equal to     | Implemented
	ASTERISK_EQUALS = "*="       // Multiply equals    | Implemented
	BANG            = "!"        // Boolean operator   | Implemented
	ASSIGN          = "="        // General assignment | Implemented
	PLUS            = "+"        // General operator   | Implemented
	MINUS           = "-"        // General operator   | Implemented
	ASTARISK        = "*"        // General operator   | Implemented
	SLASH           = "/"        // General operator   | Implemented
	LT              = "<"        // Boolean operator   | Implemented
	GT              = ">"        // Boolean operator   | Implemented
	EQ              = "=="       // Boolean operator   | Implemented
	MINUS_EQUALS    = "-="       // Minus equals       | Implemented
	NEQ             = "!="       // Boolean operator   | Implemented
	DIVEQ           = "/="       // Division operator  | Implemented
	PERIOD          = "."        // Method Call        | Implemented
	PLUS_EQUALS     = "+="       // Plus equals        | Implemented
	COMMA           = ","        // Seperation         | Implemented
	SEMICOLON       = ";"        // SemiColon          | Implemented
	COLON           = ":"        // Colon              | Implemented
	LPAREN          = "("        // Args start         | Implemented
	RPAREN          = ")"        // Args end           | Implemented
	LINE            = "|"        // Line con           | Implemented
	LBRACE          = "{"        // Open  f            | Implemented
	RBRACE          = "}"        // Close f            | Implemented
	LBRACKET        = "["        // Open               | Implemented
	RBRACKET        = "]"        // Close              | Implemented
	PLUS_PLUS       = "++"       // Plus Plus          | Not implemented
	QUESTION        = "?"        // Question que       | Not implemented
	DOTDOT          = ".."       // Range              | Not implemented
	CONTAINS        = "~="       // Contains           | Not implemented
	NOTCONTAIN      = "!~"       // Boolean operator   | Not implemented
	MINUS_MINUS     = "--"       // Minus minus        | Not Implemented
	OROR            = "||"       // Condition or or    | Not implemented
	ANDAND          = "&&"       // Boolean operator   | Not implemented
	BACKTICK        = "`"        // Backtick           | Not implemented
	POWEROF         = "**"       // General operator   | Not Implemented
	MODULO          = "%"        // General operator   | Not Implemented
	BREAK           = "BREAK"    // Break statement    | Not implemented
	DEFAULT         = "DEFAULT"  // Default keyword    | Not implemented
	MACRO           = "MACRO"    // Macro statement    | Not implemented
	REGEXP          = "REGEXP"   // Regex Type         | Not implemented
	FOR             = "FOR"      // For loop 		   | Not implemeneted
)

var keywords = map[string]Token_Type{
	"Func":     FUNCTION, // Function
	"function": FUNCTION, // Function
	"let":      LET,      // Variable declaration
	"cause":    LET,      // Variable declaration
	"allow":    LET,      // Variable declaration
	"true":     TRUE,     // Boolean true
	"false":    FALSE,    // Boolean false
	"if":       IF,       // Conditional start
	"else":     ELSE,     // Conditional alternative
	"return":   RETURN,   // Return decl
	"ret":      RETURN,   // Return decl
	"macro":    MACRO,    // Macro statement
	"const":    CONSTANT, // Constant type
	"constant": CONSTANT, // Constant type
	"switch":   SWITCH,   // Switch statement
	"sw":       SWITCH,   // Switch statement
	"case":     CASE,     // Case statement
	"cs":       CASE,     // Case statement
	"default":  DEFAULT,  // Switch alternative
	"df":       DEFAULT,  // Switch alternative
	"while":    WHILE,    // While statement
	"break":    BREAK,    // Break statement
	"foreach":  FOREACH,  // For each in statement
	"in":       IN,       // In call
	"for":      FOR,      // For loop
	"register": IMPORT,   // Register
}

///////////////////////////////////////////////////////////////////////////
////                         Lexer Types                               ////

type Lexer interface {
	NT() Token
}

type LexerStructure struct {
	CharInput string
	POS       int
	RPOS      int
	Char      byte
	Chars     []rune
	PrevTok   Token
	Prevch    byte
}

/////////////////////////////////////////////////////////////////////////////
////  						Object Types                                ////

type Type_Object string

// Data types
const (
	IntegerType     = "Integer"     // integger
	FloatType       = "Float"       // float integer
	BooleanType     = "Boolean"     // bool true or bool false
	NilType         = "Nil"         // null
	ReturnValueType = "ReturnValue" // return
	ErrorType       = "Error"       // error
	FunctionType    = "Function"    // function
	StringType      = "String"      // string
	BuiltinType     = "Builtin"     // built in
	ArrayType       = "Array"       // array
	HashType        = "Hash"        // hash
	QuoteType       = "Quote"       // quote
	MacroType       = "Macro"       // macro
	UnsignedType    = "Unsigned"    // Unsigned integer
	RegisterType    = "Registry"    // Register backend
)

type WhileCondition struct {
	Token     Token
	Condition Expression
	Body      *BlockStatement
}

type Switch struct {
	Token   Token
	Value   Expression
	Choices []*Case
}

type Case struct {
	Token Token
	Def   bool
	Expr  []Expression
	Block *BlockStatement
}

type Object interface {
	Type_Object() Type_Object
	Inspect() string
	InvokeMethod(method string, Env Environment_of_environment, args ...Object) Object
	ToInterface() interface{}
}

type HashKey struct {
	Type_Object Type_Object
	Value       uint64
}

type Hashable interface {
	HashKey() HashKey
}

type Integer struct {
	Value int64
}

type Unsigned struct {
	Value uint64
}

type Float struct {
	Value float64
}

type Boolean_Object struct {
	Value bool
}

type String struct {
	Value  string
	Offset int
}

// Offset for string reset
func (str *String) Reset() {
	str.Offset = 0
}

type Nil struct{}

type ReturnValue struct {
	Value Object
}

type Error struct {
	Message string
}

type Function struct {
	Parameters []*Ident
	Body       *BlockStatement
	Env        *Environment_of_environment
}

type BuiltinFunction func(env *Environment_of_environment, args ...Object) Object

type Constant struct {
	Token Token
	Name  *Ident
	Value Expression
}

type Builtin struct {
	Fn BuiltinFunction
}

type Array struct {
	Elements []Object
	Offset   int
}

// Offset reset for array structure
func (array *Array) Reset() {
	array.Offset = 0
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs  map[HashKey]HashPair
	Offset int
}

// Offset for hash reset
func (hash *Hash) Reset() {
	hash.Offset = 0
}

type Quote struct {
	Node
}

type Macro struct {
	Parameters []*Ident
	Body       *BlockStatement
	Env        *Environment_of_environment
}

//////////////////////////////////////////////////////////////////////////////////////
/// 									ABSTRACT SYNTAX TREE MODELS				/////

var U UserInterpretData

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	SN()
}

type Expression interface {
	Node
	EN()
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token Token
	Name  *Ident
	Value Expression
}

type AssignmentStatement struct {
	Token    Token
	Name     *Ident
	Value    Expression
	Operator string
}

type Ident struct {
	Token Token
	Value string
}

type ReturnStatement struct {
	Token       Token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token      Token
	Expression Expression
}

type UnsignedLiteral struct {
	Token Token
	Value uint64
}

type IntegerLiteral struct {
	Token Token
	Value int64
}

type FloatLiteral struct {
	Token Token
	Value float64
}

type PrefixExpression struct {
	Token    Token
	Operator string
	Right    Expression
}

type InfixExpression struct {
	Token    Token
	Left     Expression
	Operator string
	Right    Expression
}

type Boolean_AST struct {
	Token Token
	Value bool
}

type ConditionalExpression struct {
	Token       Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

type BlockStatement struct {
	Token      Token
	Statements []Statement
}

type FunctionLiteral struct {
	Token      Token
	Parameters []*Ident
	Body       *BlockStatement
}

type CallExpression struct {
	Token     Token      // the '(' token
	Function  Expression // Ident or FunctionLiteral
	Arguments []Expression
}

type StringLiteral struct {
	Token Token
	Value string
}

type ArrayLiteral struct {
	Token    Token // the '[' token
	Elements []Expression
}

type IndexExpression struct {
	Token Token // the '[' token
	Left  Expression
	Index Expression
}

type HashLiteral struct {
	Token Token // the '{' token
	Pairs map[Expression]Expression
}

type MacroLiteral struct {
	Token      Token
	Parameters []*Ident
	Body       *BlockStatement
}

//////////////////////////////////////////////////////////////////////////////////////
///                       OBJECT ENVIRONMENT MODELS AND STRUCTURES              /////

type Environment_of_environment struct {
	Store  map[string]Object
	Outer  *Environment_of_environment
	ROM    map[string]bool // Read Only mode for constants
	permit []string
}

//////////////////////////////////////////////////////////////////////////////////////
///                       			Evaluation Models                           /////

var (
	NilValue   = &Nil{}
	TrueValue  = &Boolean_Object{Value: true}
	FalseValue = &Boolean_Object{Value: false}
)

const (
	FuncNameQuote   = "quote"
	FuncNameUnquote = "unquote"
)

type ObjectCallExpression struct {
	Token  Token
	Object Expression
	Call   Expression
}

var FileCurrent FileCurrentWithinParserEnvironment

//////////////////////////////////////////////////////////////////////////////////////
/// 									PARSER MODELS							 ///

const (
	_ int = iota
	LOWEST
	ASSIGNMENT
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
	INDEX
	GTEQS
	LTEQS
	POWER
	MOD
	DOT_DOT
	REGEXP_MATCH
	TERNARY
	COND
)

var Precedences = map[Token_Type]int{
	ANDAND:          COND,
	OROR:            COND,
	EQ:              EQUALS,
	GTEQ:            GTEQS,
	LTEQ:            LTEQS,
	NEQ:             EQUALS,
	LT:              LESSGREATER,
	GT:              LESSGREATER,
	PLUS:            SUM,
	PLUS_EQUALS:     SUM,
	MINUS:           SUM,
	MINUS_EQUALS:    SUM,
	SLASH:           PRODUCT,
	ASSIGN:          ASSIGNMENT,
	POWEROF:         POWER,
	QUESTION:        TERNARY,
	ASTARISK:        PRODUCT,
	ASTERISK_EQUALS: PRODUCT,
	DIVEQ:           PRODUCT,
	LPAREN:          CALL,
	PERIOD:          CALL,
	NOTCONTAIN:      REGEXP_MATCH,
	CONTAINS:        REGEXP_MATCH,
	DOTDOT:          DOT_DOT,
	LBRACKET:        INDEX,
}

type (
	PrefixParseFn  func() Expression
	InfixParseFn   func(Expression) Expression
	PostfixParseFn func() Expression
)

type Parser struct {
	Lex             Lexer
	Errors          []string
	PreviousToken   Token
	CurrentToken    Token
	PeekToken       Token
	PrefixParseFns  map[Token_Type]PrefixParseFn
	InfixParseFns   map[Token_Type]InfixParseFn
	PostfixParseFns map[Token_Type]PostfixParseFn
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//////////// Invoking methods for the function calls to standard libraries, this modifies ENV
//////////////////////////////////////////////////////////////////////////////////////////////////

func (quote *Quote) InvokeMethod(method string, Env Environment_of_environment, args ...Object) Object {
	return nil
}

func (macro *Macro) InvokeMethod(method string, Env Environment_of_environment, args ...Object) Object {
	return nil
}

func (f *Function) InvokeMethod(method string, Env Environment_of_environment, args ...Object) Object {
	if method == "methods" {
		static := []string{"methods"}
		dynamic := Env.Names("function.")
		var names []string
		names = append(names, static...)
		for _, environ := range dynamic {
			bits := strings.Split(environ, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)
		results := make([]Object, len(names))
		for IDX, TXT := range names {
			results[IDX] = &String{Value: TXT}
		}
		return &Array{Elements: results}
	}
	return nil
}

func (Arr *Array) InvokeMethod(method string, Env Environment_of_environment, args ...Object) Object {
	if method == "len" {
		return &Integer{Value: int64(len(Arr.Elements))}
	}
	if method == "methods" {
		static := []string{"len", "methods"}
		dynamic := Env.Names("array.")
		var names []string
		names = append(names, static...)
		for _, envi := range dynamic {
			bits := strings.Split(envi, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)
		result := make([]Object, len(names))
		for idx, txt := range names {
			result[idx] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	return nil
}

func (hash *Hash) InvokeMethod(method string, Env Environment_of_environment, args ...Object) Object {
	if method == "keys" {
		ents := len(hash.Pairs)
		arrays := make([]Object, ents)
		idx := 0
		for _, entity := range hash.Pairs {
			arrays[idx] = entity.Key
			idx++
		}
		return &Array{Elements: arrays}
	}
	return nil
}

func (Bool *Boolean_Object) InvokeMethod(method string, Env Environment_of_environment, args ...Object) Object {
	if method == "methods" {
		static := []string{"mnethods"}
		dynamic := Env.Names("bool.")
		var names []string
		names = append(names, static...)
		for _, envi := range dynamic {
			bits := strings.Split(envi, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)
		result := make([]Object, len(names))
		for idx, txt := range names {
			result[idx] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	return nil
}

func (BuiltIn *Builtin) InvokeMethod(method string, Env Environment_of_environment, args ...Object) Object {
	if method == "methods" {
		names := []string{"methods"}
		result := make([]Object, len(names))
		for idx, txt := range names {
			result[idx] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	return nil
}

func (null *Nil) InvokeMethod(method string, Env Environment_of_environment, args ...Object) Object {
	return nil
}

func (Err *Error) InvokeMethod(method string, Env Environment_of_environment, args ...Object) Object {
	return nil
}

func (ret *ReturnValue) InvokeMethod(method string, Env Environment_of_environment, args ...Object) Object {
	return nil
}

func (float *Float) InvokeMethod(method string, Env Environment_of_environment, args ...Object) Object {
	if method == "methods" {
		static := []string{"methods"}
		dynamic := Env.Names("float.")
		var names []string
		names = append(names, static...)
		for _, envi := range dynamic {
			bits := strings.Split(envi, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)
		result := make([]Object, len(names))
		for idx, txt := range names {
			result[idx] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	return nil
}

func (integer *Integer) InvokeMethod(method string, Env Environment_of_environment, args ...Object) Object {
	if method == "chr" {
		return &String{Value: string(rune(integer.Value))}
	}
	if method == "methods" {
		static := []string{"chr", "methods"}
		dynamic := Env.Names("integer.")
		var names []string
		names = append(names, static...)
		for _, envi := range dynamic {
			bits := strings.Split(envi, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)
		result := make([]Object, len(names))
		for idx, txt := range names {
			result[idx] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	return nil
}

func (Str *String) InvokeMethod(method string, Env Environment_of_environment, args ...Object) Object {
	if method == "len" {
		return &Integer{Value: int64(utf8.RuneCountInString(Str.Value))}
	}
	if method == "methods" {
		static := []string{"len", "methods", "ord", "to_i", "to_f"}
		dynamic := Env.Names("string.")
		var names []string
		names = append(names, static...)
		for _, envi := range dynamic {
			bits := strings.Split(envi, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)
		result := make([]Object, len(names))
		for Idx, txt := range names {
			result[Idx] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	if method == "ord" {
		return &Integer{Value: int64(Str.Value[0])}
	}
	if method == "to_i" {
		idx, x := strconv.ParseInt(Str.Value, 0, 64)
		if x != nil {
			idx = 0
		}
		return &Integer{Value: int64(idx)}
	}
	if method == "to_f" {
		idx, x := strconv.ParseFloat(Str.Value, 64)
		if x != nil {
			idx = 0
		}
		return &Float{Value: idx}
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////// ENVIRONMENT MODIFIER
//////////////////////////////////////////////////////////////////////////////////////////////////

type Iterable interface {
	Reset()
	Next() (Object, Object, bool)
}

func (env *Environment_of_environment) Names(prefix string) []string {
	var ret []string
	for key := range env.Store {
		if strings.HasPrefix(key, prefix) {
			ret = append(ret, key)
		}
		if strings.HasPrefix(key, "object.") {
			ret = append(ret, key)
		}
	}
	return ret
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//////////// This is all methods for to interface for values and standard calls
//////////////////////////////////////////////////////////////////////////////////////////////////

func (quote *Quote) ToInterface() interface{} {
	return "<QUOTE>"
}

func (macro *Macro) ToInterface() interface{} {
	return "<MACRO>"
}

func (Arr *Array) ToInterface() interface{} {
	return "<ARRAY>"
}

func (hash *Hash) ToInterface() interface{} {
	return "<HASH>"
}

func (null *Nil) ToInterface() interface{} {
	return "<NULL>"
}

func (Func *Function) ToInterface() interface{} {
	return "<FUNCTION>"
}

func (Err *Error) ToInterface() interface{} {
	return "<ERROR>"
}

func (BuiltIn *Builtin) ToInterface() interface{} {
	return "<BUILT-IN-FUNCTION>"
}

func (Ret *ReturnValue) ToInterface() interface{} {
	return "<RETURN_VALUE>"
}

func (float *Float) ToInterface() interface{} {
	return float.Value
}

func (Int *Integer) ToInterface() interface{} {
	return Int.Value
}

func (Str *String) ToInterface() interface{} {
	return Str.Value
}

func (Boolean *Boolean_Object) ToInterface() interface{} {
	return Boolean.Value
}

////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// NEXT METHODS FOR STRING, HASH, ARRAY
////////////////////////////////////////////////////////////////////////////////////////

func (Arr *Array) Next() (Object, Object, bool) {
	if Arr.Offset < len(Arr.Elements) {
		Arr.Offset++
		element := Arr.Elements[Arr.Offset-1]
		return element, &Integer{Value: int64(Arr.Offset - 1)}, true
	}
	return nil, &Integer{Value: 0}, false
}

func (hash *Hash) Next() (Object, Object, bool) {
	if hash.Offset < len(hash.Pairs) {
		idx := 0
		for _, pair := range hash.Pairs {
			if hash.Offset == idx {
				hash.Offset++
				return pair.Key, pair.Value, true
			}
			idx++
		}
	}
	return nil, &Integer{Value: 0}, false
}

func (str *String) Next() (Object, Object, bool) {
	if str.Offset < utf8.RuneCountInString(str.Value) {
		str.Offset++
		chars := []rune(str.Value)
		value := &String{Value: string(chars[str.Offset-1])}
		return value, &Integer{Value: int64(str.Offset - 1)}, true
	}
	return nil, &Integer{Value: 0}, false
}
