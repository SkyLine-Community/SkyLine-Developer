package SkyLine_Backend

import "fmt"

func (parser *Parser) ParserErrors() []string          { return parser.Errors }
func (parser *Parser) PeekTokenIs(typ Token_Type) bool { return parser.PeekToken.Token_Type == typ }

func (parser *Parser) CurrentTokenIs(typ Token_Type) bool {
	return parser.CurrentToken.Token_Type == typ
}

func (parser *Parser) NT() {
	parser.CurrentToken = parser.PeekToken
	parser.PeekToken = parser.Lex.NT()
}

func (parser *Parser) PeekError(typ Token_Type) {
	msg := Map_Parser[ERROR_DURING_PEEK_IN_PARSER](fmt.Sprint(typ), string(parser.PeekToken.Token_Type)).Message
	parser.Errors = append(parser.Errors, msg)
}

func (parser *Parser) ExpectPeek(typ Token_Type) bool {
	if parser.PeekTokenIs(typ) {
		parser.NT()
		return true
	}
	parser.PeekError(typ)
	return false
}

func (parser *Parser) ParseProgram() *Program {
	program := &Program{
		Statements: []Statement{},
	}

	for !parser.CurrentTokenIs(EOF) {
		stmt := parser.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		parser.NT()
	}

	return program
}
