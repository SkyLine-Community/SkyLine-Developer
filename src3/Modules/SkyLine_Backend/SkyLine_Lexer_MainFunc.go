package SkyLine

func (lex *LexerStructure) NextToken() Token {
	lex.ConsumeWhiteSpace()
	// Hybrid note system
	if lex.Char == '/' && lex.Peek() == '/' {
		lex.ConsumeComment()
	}
	if lex.Char == '!' && lex.Peek() == '!' {
		lex.ConsumeComment()
	}
	if lex.Char == '#' {
		lex.ConsumeComment()
	}

	var tok Token
	switch lex.Char {
	case '=':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: EQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = newToken(ASSIGN, lex.Char)
		}
	case '!':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: NEQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = newToken(BANG, lex.Char)
		}
	case ';':
		tok = newToken(SEMICOLON, lex.Char)
	case ':':
		tok = newToken(COLON, lex.Char)
	case '(':
		tok = newToken(LPAREN, lex.Char)
	case '|':
		tok = newToken(LINE, lex.Char)
	case ')':
		tok = newToken(RPAREN, lex.Char)
	case ',':
		tok = newToken(COMMA, lex.Char)
	case '+':
		tok = newToken(PLUS, lex.Char)
	case '-':
		tok = newToken(MINUS, lex.Char)
	case '*':
		tok = newToken(ASTARISK, lex.Char)
	case '/':
		tok = newToken(SLASH, lex.Char)
	case '<':
		tok = newToken(LT, lex.Char)
	case '>':
		tok = newToken(GT, lex.Char)
	case '{':
		tok = newToken(LBRACE, lex.Char)
	case '}':
		tok = newToken(RBRACE, lex.Char)
	case '[':
		tok = newToken(LBRACKET, lex.Char)
	case ']':
		tok = newToken(RBRACKET, lex.Char)
	case '"':
		tok.Token_Type = STRING
		tok.Literal = lex.ReadString()
	case 0:
		tok.Literal = ""
		tok.Token_Type = EOF
	default:
		if CharIsDigit(lex.Char) {
			return lex.ReadIntToken()
		}

		if CharIsLetter(lex.Char) {
			tok.Literal = lex.ReadIdentifier()
			tok.Token_Type = LookupIdentifier(tok.Literal)
			return tok
		}

		tok = newToken(ILLEGAL, lex.Char)
	}

	lex.ReadChar()
	return tok
}
