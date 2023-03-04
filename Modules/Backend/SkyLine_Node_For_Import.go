package SkyLine_Backend

import (
	"bytes"
	"fmt"
)

type RegisterValue struct {
	Value Object
}

type Register struct {
	Token       Token
	ImportValue Expression
}

const (
	ImportingType = "Import"
)

func (importvalue *RegisterValue) Type_Object() Type_Object { return ImportingType }
func (importvalue *Register) SN()                           {}                                   // Statement Node | Allow import
func (importvalue *Register) TokenLiteral() string          { return importvalue.Token.Literal } // Token Literal | Returns literal of importing list
func (importvalue *RegisterValue) Inspect() string          { return importvalue.Value.Inspect() }

func (importvalue *Register) String() string {
	var Out bytes.Buffer
	Out.WriteString("register (")
	Out.WriteString(importvalue.ImportValue.String())
	Out.WriteString(")")
	return Out.String()
}

func (importvalue *RegisterValue) InvokeMethod(method string, Env Environment_of_environment, args ...Object) Object {
	return nil
}

func (importvalue *RegisterValue) ToInterface() interface{} {
	return "<REGISTRY>"
}

func (parser *Parser) ParseImportStatement() *Register {
	statement := &Register{Token: parser.CurrentToken}
	if !parser.PeekTokenIs("(") {
		fmt.Println("Missing ( in import statement")
	}
	parser.NT()
	statement.ImportValue = parser.parseExpression(LOWEST)
	for parser.PeekTokenIs(SEMICOLON) {
		parser.NT()
	}
	return statement
}
