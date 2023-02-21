package SkyLine_Backend

import (
	"bytes"
	"fmt"
)

type Registry struct {
	Token Token
	Value Expression
}

type Registration_Entry struct {
	Value Object
}

// Methods for registry - Type Registry
func (Register *Registry) InvokeMethod(method string, Env Environment_of_environment, args ...Object) Object {
	return nil
}

func (Register *Registry) ToInterface() interface{} {
	return "<REGISTER>"
}

func (Register *Registry) SN() {}

func (Register *Registry) String() string {
	var Out bytes.Buffer
	Out.WriteString("register(")
	Out.WriteString(Register.Value.String())
	Out.WriteString(")")
	return Out.String()
}

func (parser *Parser) ParseRegister() Expression {
	stmt := &Registry{Token: parser.CurrentToken}
	if !parser.PeekTokenIs("(") {
		fmt.Println("missing ( in register")
	}
	parser.NT()
	stmt.Value = parser.parseExpression(LOWEST)
	return stmt
}
