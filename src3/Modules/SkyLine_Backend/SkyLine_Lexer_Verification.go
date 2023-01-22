package SkyLine

// These functions verify the token types

func CharIsLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func CharIsDigit(char byte) bool { return '0' <= char && char <= '9' }

func newToken(TT Token_Type, ch byte) Token {
	return Token{
		Token_Type: TT,
		Literal:    string(ch),
	}
}
