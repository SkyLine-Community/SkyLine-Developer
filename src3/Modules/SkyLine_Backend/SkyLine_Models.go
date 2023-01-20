package SkyLine

///////////////////////////////////////////////////////////////////////////
//// 						Tokenizer types                          ////

type Token_Type string

type Token struct {
	Token_Type Token_Type
	Literal    string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	FLOAT     = "FLOAT"
	STRING    = "STRING"
	BANG      = "!"
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	ASTARISK  = "*"
	SLASH     = "/"
	LT        = "<"
	GT        = ">"
	EQ        = "=="
	NEQ       = "!="
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LINE      = "|"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
	MACRO     = "MACRO"
	CARRIER   = "CARRIER"
	IMPORT    = "IMPORT"
)

var keywords = map[string]Token_Type{
	"Func":     FUNCTION,
	"function": FUNCTION,
	"let":      LET,
	"cause":    LET,
	"allow":    LET,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"ret":      RETURN,
	"macro":    MACRO,
	"carry":    CARRIER,
	"import":   IMPORT,
	"require":  IMPORT,
	"include":  IMPORT,
}

///////////////////////////////////////////////////////////////////////////
////                         Lexer Types                               ////

type Lexer interface {
	NextToken() Token
}

type LexerStructure struct {
	CharInput string
	POS       int
	RPOS      int
	Char      byte
}

/////////////////////////////////////////////////////////////////////////////
////  						Object Types                                ////

type Type_Object string

const (
	IntegerType     Type_Object = "Integer"
	FloatType                   = "Float"
	BooleanType                 = "Boolean"
	NilType                     = "Nil"
	ReturnValueType             = "ReturnValue"
	ErrorType                   = "Error"
	FunctionType                = "Function"
	StringType                  = "String"
	BuiltinType                 = "Builtin"
	ArrayType                   = "Array"
	HashType                    = "Hash"
	QuoteType                   = "Quote"
	MacroType                   = "Macro"
	CarrierType                 = "Carrier"
	ImportingType               = "Importing"
)

type Object interface {
	Type_Object() Type_Object
	Inspect() string
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

type Float struct {
	Value float64
}

type Boolean_Object struct {
	Value bool
}

type String struct {
	Value string
}

type Nil struct{}

type CarrierValue struct {
	Value Object
}

type ImporterValue struct {
	Value Object
}

type ReturnValue struct {
	Value Object
}

type Error struct {
	Message string
}

type Function struct {
	Parameters []*Ident
	Body       *BlockStatement
	Env        Environment
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

type Array struct {
	Elements []Object
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

type Quote struct {
	Node
}

type Macro struct {
	Parameters []*Ident
	Body       *BlockStatement
	Env        Environment
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

type Ident struct {
	Token Token
	Value string
}

type ReturnStatement struct {
	Token       Token
	ReturnValue Expression
}

// Expression node
type Carrier struct {
	Token        Token
	CarrierValue Expression
}

// Expression node for importing files
type Import struct {
	Token       Token
	ImportValue Expression
}

type ExpressionStatement struct {
	Token      Token
	Expression Expression
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

type Environment interface {
	Get(name string) (Object, bool)
	Set(name string, val Object) Object
}

type Environment_of_environment struct {
	Store map[string]Object
	Outer Environment
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

var FileCurrent FileCurrentWithinParserEnvironment

//////////////////////////////////////////////////////////////////////////////////////
/// 									PARSER MODELS							 ///

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunc(X)
	INDEX       // array[index]
)

var Precedences = map[Token_Type]int{
	EQ:       EQUALS,
	NEQ:      EQUALS,
	LT:       LESSGREATER,
	GT:       LESSGREATER,
	PLUS:     SUM,
	MINUS:    SUM,
	SLASH:    PRODUCT,
	ASTARISK: PRODUCT,
	LPAREN:   CALL,
	LBRACKET: INDEX,
}

type (
	PrefixParseFn func() Expression
	InfixParseFn  func(Expression) Expression
)

type Parser struct {
	Lex            Lexer
	Errors         []string
	CurrentToken   Token
	PeekToken      Token
	PrefixParseFns map[Token_Type]PrefixParseFn
	InfixParseFns  map[Token_Type]InfixParseFn
}
