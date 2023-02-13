package SkyLine

import "strings"

func (lex *LexerStructure) ReadIdentifier() string { return lex.read(CharIsLetter) }
func (lex *LexerStructure) ReadNumber() string {
	NewCharLst := ""
	InterpreterAccepts := "0123456789"
	if lex.Char == '0' && lex.Peek() == 'x' {
		InterpreterAccepts = "0x123456789abcdefABCDEF"
	}
	if lex.Char == '0' && lex.Peek() == 'b' {
		InterpreterAccepts = "b01"
	}
	for strings.Contains(InterpreterAccepts, string(lex.Char)) {
		NewCharLst += string(lex.Char)
		lex.ReadChar()
	}
	return NewCharLst

}

func (lex *LexerStructure) ReadChar() {
	if lex.RPOS >= len(lex.CharInput) {
		lex.Char = 0
	} else {
		lex.Char = lex.CharInput[lex.RPOS]
	}
	lex.POS = lex.RPOS
	lex.RPOS++
}

func (lex *LexerStructure) ReadString() string {
	outform := ""
	for {
		lex.ReadChar()
		if lex.Char == '"' || lex.Char == 0 {
			break
		}
		if lex.Char == '\\' {
			lex.ReadChar()
			switch lex.Char {
			case 'n':
				lex.Char = '\n'
			case 'r':
				lex.Char = '\r'
			case 't':
				lex.Char = '\t'
			case 'f':
				lex.Char = '\f'
			case 'v':
				lex.Char = '\v'
			case '\\':
				lex.Char = '\\'
			case '"':
				lex.Char = '"'
			}
		}
		outform = outform + string(lex.Char)
	}
	return outform
}

func (lex *LexerStructure) read(checkFn func(byte) bool) string {
	position := lex.POS
	for checkFn(lex.Char) {
		lex.ReadChar()
	}
	return lex.CharInput[position:lex.POS]
}

func (lex *LexerStructure) ReadIntToken() Token {
	intPart := lex.ReadNumber()
	if lex.Char != '.' {
		return Token{
			Token_Type: INT,
			Literal:    intPart,
		}
	}

	lex.ReadChar()
	fracPart := lex.ReadNumber()
	return Token{
		Token_Type: FLOAT,
		Literal:    intPart + "." + fracPart,
	}
}

func LexNew(input string) *LexerStructure {
	lex := &LexerStructure{CharInput: input}
	lex.ReadChar()
	return lex
}

func (lex *LexerStructure) Peek() byte {
	if lex.RPOS >= len(lex.CharInput) {
		return 0
	}
	return lex.CharInput[lex.RPOS]
}
