package SkyLine_Backend

func (parser *Parser) NextLoadFaultToken() {
	parser.PreviousToken = parser.CurrentToken
	parser.CurrentToken = parser.PeekToken
	parser.PeekToken = parser.Lex.NT()
}
