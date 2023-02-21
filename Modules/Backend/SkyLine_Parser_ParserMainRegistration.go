package SkyLine_Backend

func New_Parser(l Lexer) *Parser {
	parser := &Parser{
		Lex:    l,
		Errors: []string{},
	}
	parser.NextLoadFaultToken()
	parser.NextLoadFaultToken()

	parser.PrefixParseFns = map[Token_Type]PrefixParseFn{
		UNSIGNED: parser.ParseUnsignedLiteral,
		IDENT:    parser.parseIdent,
		INT:      parser.parseIntegerLiteral,
		FLOAT:    parser.parseFloatLiteral,
		BANG:     parser.parsePrefixExpression,
		MINUS:    parser.parsePrefixExpression,
		TRUE:     parser.parseBoolean,
		FALSE:    parser.parseBoolean,
		LPAREN:   parser.parseGroupedExpression,
		IF:       parser.parseIfExpression,
		FUNCTION: parser.parseFunctionLiteral,
		STRING:   parser.parseStringLiteral,
		LBRACKET: parser.parseArrayLiteral,
		LBRACE:   parser.parseHashLiteral,
		LINE:     parser.ParseGroupImportExpression,
		MACRO:    parser.parseMacroLiteral,
		SWITCH:   parser.ParseSwitchCaseStatement,
		WHILE:    parser.ParseWhileExpression,
		REGISTER: parser.parseGroupedExpression,
	}

	parser.InfixParseFns = map[Token_Type]InfixParseFn{
		PLUS:            parser.parseInfixExpression,
		MINUS:           parser.parseInfixExpression,
		ASTARISK:        parser.parseInfixExpression,
		SLASH:           parser.parseInfixExpression,
		EQ:              parser.parseInfixExpression,
		NEQ:             parser.parseInfixExpression,
		LT:              parser.parseInfixExpression,
		GT:              parser.parseInfixExpression,
		LPAREN:          parser.parseCallExpression,
		LBRACKET:        parser.parseIndexExpression,
		GTEQ:            parser.parseInfixExpression,
		LTEQ:            parser.parseInfixExpression,
		PLUS_EQUALS:     parser.parseAssignmentStatement,
		DIVEQ:           parser.parseAssignmentStatement,
		MINUS_EQUALS:    parser.parseAssignmentStatement,
		ASTERISK_EQUALS: parser.parseAssignmentStatement,
		ASSIGN:          parser.parseAssignmentStatement,
		ANDAND:          parser.parseInfixExpression,
		OROR:            parser.parseInfixExpression,
		POWEROF:         parser.parseInfixExpression,
		PERIOD:          parser.parseMethodCallExpression,
	}

	// Read two tokens, so curToken and peekToken are both set

	return parser
}
