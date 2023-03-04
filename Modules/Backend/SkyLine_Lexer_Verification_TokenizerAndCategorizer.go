package SkyLine_Backend

func (lex *LexerStructure) NT() Token {
	lex.ConsumeWhiteSpace()
	if lex.Char == '/' && lex.Peek() == '/' {
		lex.ConsumeComment()
		return lex.NT()
	}
	if lex.Char == '/' && lex.Peek() == '*' {
		lex.ConsumeMultiLineComment()
		return lex.NT()
	}
	if lex.Char == '!' && lex.Peek() == '-' {
		lex.ConsumeMultiLineComment()
		return lex.NT()
	}

	var tok Token
	switch lex.Char {
	case '.':
		if lex.Peek() == '.' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: DOTDOT,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = newToken(PERIOD, lex.Char)
		}
	case '&':
		if lex.Peek() == '&' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: ANDAND,
				Literal:    string(ch) + string(lex.Char),
			}
		}
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
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: ASSIGN,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = newToken(COLON, lex.Char)
		}
	case '(':
		tok = newToken(LPAREN, lex.Char)
	case '|':
		if lex.Peek() == '|' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: OROR,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = newToken(LINE, lex.Char)
		}
	case ')':
		tok = newToken(RPAREN, lex.Char)
	case ',':
		tok = newToken(COMMA, lex.Char)
	case '+':
		tok = newToken(PLUS, lex.Char)
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: PLUS_EQUALS,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = newToken(PLUS, lex.Char)
		}
	case '-':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: MINUS_EQUALS,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = newToken(MINUS, lex.Char)
		}
	case '*':
		if lex.Peek() == '*' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: POWEROF,
				Literal:    string(ch) + string(lex.Char),
			}
		} else if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: ASTERISK_EQUALS,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = newToken(ASTARISK, lex.Char)
		}
	case '%':
		tok = newToken(MODULO, lex.Char)
	case '/':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: DIVEQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = newToken(SLASH, lex.Char)
		}
	case '<':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: LTEQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = newToken(LT, lex.Char)
		}
	case '>':
		if lex.Peek() == '=' {
			cha := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: GTEQ,
				Literal:    string(cha) + string(lex.Char),
			}
		} else {
			tok = newToken(GT, lex.Char)
		}
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
			lex.PrevTok = tok
			return tok
		}

		tok = newToken(ILLEGAL, lex.Char)
	}
	lex.ReadChar()
	return tok
}
