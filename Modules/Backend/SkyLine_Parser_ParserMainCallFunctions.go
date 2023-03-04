package SkyLine_Backend

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (parser *Parser) parseStatement() Statement {
	switch parser.CurrentToken.Token_Type {
	case IMPORT:
		return parser.ParseImportStatement()
	case LET:
		return parser.ParserCreateAssignment()
	case RETURN:
		return parser.parseReturnStatement()
	case CONSTANT:
		return parser.ParseConstants()
	default:
		return parser.parseExpressionStatement()
	}
}

func (parser *Parser) ParseSwitchCaseStatement() Expression {
	EXP := &Switch{Token: parser.CurrentToken}
	EXP.Value = parser.ParseArgumentParens()
	if EXP.Value == nil {
		return nil
	}
	if !parser.ExpectPeek(LBRACE) {
		return nil
	}
	parser.NT()
	for !parser.CurrentTokenIs(RBRACE) {
		if parser.CurrentTokenIs(EOF) {
			parser.Errors = append(parser.Errors, "unterminated switch statement")
		}
		exp := &Case{Token: parser.CurrentToken}
		if parser.CurrentTokenIs(DEFAULT) {
			exp.Def = true

		} else if parser.CurrentTokenIs(CASE) {

			parser.NT()
			if parser.CurrentTokenIs(DEFAULT) {
				exp.Def = true
			} else {
				exp.Expr = append(exp.Expr, parser.parseExpression(LOWEST))
				for parser.PeekTokenIs(COMMA) {
					parser.NT()
					parser.NT()
					exp.Expr = append(exp.Expr, parser.parseExpression(LOWEST))
				}
			}
		} else {
			parser.Errors = append(parser.Errors, fmt.Sprintf("expected case|default, got %s >>> %s ", parser.CurrentToken.Token_Type, parser.CurrentToken))
			return nil
		}

		if !parser.ExpectPeek(LBRACE) {

			msg := fmt.Sprintf("expected token to be '{', got %s instead", parser.CurrentToken.Token_Type)
			parser.Errors = append(parser.Errors, msg)
			fmt.Printf("error\n")
			return nil
		}

		// parse the block
		exp.Block = parser.parseBlockStatement()

		if !parser.CurrentTokenIs(RBRACE) {
			msg := fmt.Sprintf("Syntax Error: expected token to be '}', got %s instead", parser.CurrentToken.Token_Type)
			parser.Errors = append(parser.Errors, msg)
			fmt.Printf("error\n")
			return nil

		}
		parser.NT()
		EXP.Choices = append(EXP.Choices, exp)
	}
	count := 0
	for _, c := range EXP.Choices {
		if c.Def {
			count++
		}
	}
	if count > 1 {
		msg := "A switch-statement should only have one default block"
		parser.Errors = append(parser.Errors, msg)
		return nil

	}
	return EXP

}

func (parser *Parser) ParseArgumentParens() Expression {
	if !parser.ExpectPeek(LPAREN) {
		parser.Errors = append(parser.Errors, fmt.Sprintf("Unexpected token | %s | I need %s for this argument list ", parser.CurrentToken.Literal, LPAREN))
		return nil
	}
	parser.NT()
	exp := parser.parseExpression(LOWEST)
	if exp == nil {
		return nil
	}
	if !parser.ExpectPeek(RPAREN) {
		parser.Errors = append(parser.Errors, fmt.Sprintf("Unexpected token | %s|  I need %s for this statement ", parser.CurrentToken.Literal, RPAREN))
		return nil
	}
	return exp
}

func (parser *Parser) ParserCreateAssignment() *LetStatement {
	stmt := &LetStatement{Token: parser.CurrentToken}
	if !parser.ExpectPeek(IDENT) {
		return nil
	}
	stmt.Name = &Ident{Token: parser.CurrentToken, Value: parser.CurrentToken.Literal}
	if !parser.ExpectPeek(ASSIGN) {
		return nil
	}
	parser.NT()
	stmt.Value = parser.parseExpression(LOWEST)
	for !parser.CurrentTokenIs(SEMICOLON) {
		if parser.CurrentTokenIs(EOF) {
			parser.Errors = append(parser.Errors, "Unterminated allow or assignment statement ASSIGN/LET/ALLOW")
		}
		parser.NT()
	}
	return stmt

}

func (parser *Parser) parseAssignmentStatement(name Expression) Expression {
	stmt := &AssignmentStatement{Token: parser.CurrentToken}
	if StatementName, ok := name.(*Ident); ok {
		stmt.Name = StatementName
	} else {
		parser.Errors = append(parser.Errors, fmt.Sprintf("Expected assignment token before operator to be an IDENTIFIER not %s", name.TokenLiteral()))
	}
	opperand := parser.CurrentToken
	parser.NT()
	switch opperand.Token_Type {
	case PLUS_EQUALS:
		stmt.Operator = "+="
	case MINUS_EQUALS:
		stmt.Operator = "-="
	case DIVEQ:
		stmt.Operator = "/="
	case ASTERISK_EQUALS:
		stmt.Operator = "*="
	default:
		stmt.Operator = "="
	}

	stmt.Value = parser.parseExpression(LOWEST)
	return stmt
}

func (parser *Parser) ParseConstants() *Constant {
	statement := &Constant{Token: parser.CurrentToken}
	if !parser.ExpectPeek(IDENT) {
		return nil
	}
	statement.Name = &Ident{Token: parser.CurrentToken, Value: parser.CurrentToken.Literal}
	if !parser.ExpectPeek(ASSIGN) {
		return nil
	}
	parser.NT()
	statement.Value = parser.parseExpression(LOWEST)
	for !parser.CurrentTokenIs(SEMICOLON) {
		if parser.CurrentTokenIs(EOF) {
			parser.Errors = append(parser.Errors, "unterminated CONSTANT")
			return nil
		}
		parser.NT()
	}
	return statement
}

func (parser *Parser) parseReturnStatement() *ReturnStatement {
	stmt := &ReturnStatement{
		Token: parser.CurrentToken,
	}
	parser.NT()

	stmt.ReturnValue = parser.parseExpression(LOWEST)

	for parser.PeekTokenIs(SEMICOLON) {
		parser.NT()
	}

	return stmt
}

func (parser *Parser) parseExpressionStatement() *ExpressionStatement {
	stmt := &ExpressionStatement{
		Token:      parser.CurrentToken,
		Expression: parser.parseExpression(LOWEST),
	}

	if parser.PeekTokenIs(SEMICOLON) {
		parser.NT()
	}

	return stmt
}

func (parser *Parser) parseExpression(precedence int) Expression {
	prefix := parser.PrefixParseFns[parser.CurrentToken.Token_Type]

	if prefix == nil {
		msg := "No prefix parse function found for token " + parser.CurrentToken.Literal + " Could not locate parse function for the parsed token"
		parser.Errors = append(parser.Errors, msg)
		return nil
	}

	leftExp := prefix()

	for !parser.CurrentTokenIs(SEMICOLON) && precedence < parser.peekPrecedence() {
		infix := parser.InfixParseFns[parser.PeekToken.Token_Type]
		if infix == nil {
			return leftExp
		}

		parser.NT()

		leftExp = infix(leftExp)
	}
	return leftExp
}

func (parser *Parser) parseIdent() Expression {
	return &Ident{
		Token: parser.CurrentToken,
		Value: parser.CurrentToken.Literal,
	}
}

func sp() {
	fmt.Println()
}

func DrawTable(rowdata [][]string, columndata []string) {
	colwidth := make([]int, len(columndata))
	for k, col := range columndata {
		colwidth[k] = len(col)
	}
	for _, r := range rowdata {
		for i, table_cell := range r {
			tablecwidth := len(table_cell)
			if tablecwidth > colwidth[i] {
				colwidth[i] = tablecwidth
			}
		}
	}
	for i, col := range columndata {
		fmt.Printf("┃ %-*s ", colwidth[i], col)
	}
	sp()
	for _, w := range colwidth {
		fmt.Printf("┃ %s ", strings.Repeat("━", w))
	}
	sp()
	for _, r := range rowdata {
		for k, c := range r {
			fmt.Printf("┃ %-*s ", colwidth[k], c)
		}
		sp()
	}
}

func (parser *Parser) ParseUnsignedLiteral() Expression {
	lit := &UnsignedLiteral{Token: parser.CurrentToken}
	value, x := strconv.ParseUint(parser.CurrentToken.Literal, 0, 64)
	if x != nil {
		fmt.Println("Could not parse unsigned literal -> ", x)
	}
	lit.Value = value
	return lit
}

//TODO: Remake this error system for 90%
func (parser *Parser) parseIntegerLiteral() Expression {
	lit := &IntegerLiteral{Token: parser.CurrentToken}
	var value int64
	var x error
	if strings.HasPrefix(parser.CurrentToken.Literal, "0b") {
		value, x = strconv.ParseInt(parser.CurrentToken.Literal[2:], 2, 64)
	} else if strings.HasPrefix(parser.CurrentToken.Literal, "0x") {
		value, x = strconv.ParseInt(parser.CurrentToken.Literal[2:], 16, 64)
	} else {
		value, x = strconv.ParseInt(parser.CurrentToken.Literal, 0, 64)
	}
	if x != nil {
		fmt.Println(Map_Parser[ERROR_TYPE_INTEGRITY_PARSE_INTEGER_ERROR](parser.CurrentToken.Literal))
	}
	lit.Value = value
	return lit
}

func (parser *Parser) parseFloatLiteral() Expression {
	val, err := strconv.ParseFloat(parser.CurrentToken.Literal, 64)
	if err != nil {
		msg := "Could not parse (  " + fmt.Sprint(parser.CurrentToken.Literal) + " ) as FLOAT due to -> " + fmt.Sprint(err)
		parser.Errors = append(parser.Errors, msg)
		return nil
	}

	return &FloatLiteral{
		Token: parser.CurrentToken,
		Value: val,
	}
}

func (parser *Parser) parsePrefixExpression() Expression {
	expr := &PrefixExpression{
		Token:    parser.CurrentToken,
		Operator: parser.CurrentToken.Literal,
	}

	parser.NT()

	expr.Right = parser.parseExpression(PREFIX)
	return expr
}

func (parser *Parser) peekPrecedence() int {
	if p, ok := Precedences[parser.PeekToken.Token_Type]; ok {
		return p
	}
	return LOWEST
}

func (parser *Parser) curPrecedence() int {
	if p, ok := Precedences[parser.CurrentToken.Token_Type]; ok {
		return p
	}
	return LOWEST
}

func (parser *Parser) parseInfixExpression(left Expression) Expression {
	expr := &InfixExpression{
		Token:    parser.CurrentToken,
		Operator: parser.CurrentToken.Literal,
		Left:     left,
	}

	prec := parser.curPrecedence()

	parser.NT()

	expr.Right = parser.parseExpression(prec)
	return expr
}

func (parser *Parser) parseBoolean() Expression {
	return &Boolean_AST{
		Token: parser.CurrentToken,
		Value: parser.CurrentTokenIs(TRUE),
	}
}

func (parser *Parser) ParseGroupImportExpression() Expression {
	parser.NT()

	expr := parser.parseExpression(LOWEST)

	if !parser.ExpectPeek(LINE) {
		return nil
	}

	return expr
}

func (parser *Parser) parseGroupedExpression() Expression {
	parser.NT()

	expr := parser.parseExpression(LOWEST)

	if !parser.ExpectPeek(RPAREN) {
		return nil
	}

	return expr
}

func (parser *Parser) parseIfExpression() Expression {
	expr := &ConditionalExpression{Token: parser.CurrentToken}

	parser.NT()
	expr.Condition = parser.parseExpression(LOWEST)
	if !parser.ExpectPeek(LBRACE) {
		return nil
	}

	expr.Consequence = parser.parseBlockStatement()
	if expr.Consequence == nil {
		parser.Errors = append(parser.Errors, "unexpected nil expression")
		return nil
	}
	if parser.PeekTokenIs(ELSE) {
		parser.NT()

		if parser.PeekTokenIs(IF) {
			parser.NT()
			expr.Alternative = &BlockStatement{
				Statements: []Statement{
					&ExpressionStatement{
						Expression: parser.parseIfExpression(),
					},
				},
			}
			return expr
		}

		// parse else

		if !parser.ExpectPeek(LBRACE) {
			msg := fmt.Sprintf("expected '{' but got %s", parser.CurrentToken.Literal)
			parser.Errors = append(parser.Errors, msg)
			return nil
		}

		expr.Alternative = parser.parseBlockStatement()
		if expr.Alternative == nil {
			parser.Errors = append(parser.Errors, "unexpected nil expression")
			return nil
		}
	}

	return expr
}

func (parser *Parser) parseBlockStatement() *BlockStatement {
	block := &BlockStatement{
		Token:      parser.CurrentToken,
		Statements: []Statement{},
	}

	parser.NT()
	for !parser.CurrentTokenIs(RBRACE) && !parser.CurrentTokenIs(EOF) {
		stmt := parser.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}

		parser.NT()
	}
	if parser.CurrentTokenIs(RBRACE) && !parser.PeekTokenIs(SEMICOLON) {
		fmt.Println("Error: Could not evaluate block statement further, missing semicolon after end of function in '}' ")
		os.Exit(1)
	}

	return block
}

func (parser *Parser) parseFunctionLiteral() Expression {
	lit := &FunctionLiteral{Token: parser.CurrentToken}

	if !parser.ExpectPeek(LPAREN) {
		return nil
	}

	lit.Parameters = parser.parseFunctionParameters()

	if !parser.ExpectPeek(LBRACE) {
		return nil
	}

	lit.Body = parser.parseBlockStatement()

	return lit
}

func (parser *Parser) parseFunctionParameters() []*Ident {
	idents := []*Ident{}

	if parser.PeekTokenIs(RPAREN) {
		parser.NT()
		return idents
	}

	parser.NT()

	ident := &Ident{
		Token: parser.CurrentToken,
		Value: parser.CurrentToken.Literal,
	}
	idents = append(idents, ident)

	for parser.PeekTokenIs(COMMA) || parser.PeekTokenIs(COLON) {
		parser.NT()
		parser.NT()
		ident := &Ident{
			Token: parser.CurrentToken,
			Value: parser.CurrentToken.Literal,
		}
		idents = append(idents, ident)
	}

	if !parser.ExpectPeek(RPAREN) {
		return nil
	}

	return idents
}

func (parser *Parser) parseExpressionList(end Token_Type) []Expression {
	list := make([]Expression, 0)

	if parser.PeekTokenIs(end) {
		parser.NT()
		return list
	}

	parser.NT()
	list = append(list, parser.parseExpression(LOWEST))

	for parser.PeekTokenIs(COMMA) {
		parser.NT()
		parser.NT()
		list = append(list, parser.parseExpression(LOWEST))
	}

	if !parser.ExpectPeek(end) {
		return nil
	}

	return list
}

func (parser *Parser) parseCallExpression(function Expression) Expression {
	return &CallExpression{
		Token:     parser.CurrentToken,
		Function:  function,
		Arguments: parser.parseExpressionList(RPAREN),
	}
}

func (parser *Parser) parseStringLiteral() Expression {
	return &StringLiteral{
		Token: parser.CurrentToken,
		Value: parser.CurrentToken.Literal,
	}
}

func (parser *Parser) parseArrayLiteral() Expression {
	return &ArrayLiteral{
		Token:    parser.CurrentToken,
		Elements: parser.parseExpressionList(RBRACKET),
	}
}

func (parser *Parser) parseIndexExpression(left Expression) Expression {
	expr := &IndexExpression{
		Token: parser.CurrentToken,
		Left:  left,
	}

	parser.NT()
	expr.Index = parser.parseExpression(LOWEST)

	if !parser.ExpectPeek(RBRACKET) {
		return nil
	}

	return expr
}

func (parser *Parser) parseHashLiteral() Expression {
	hash := &HashLiteral{
		Token: parser.CurrentToken,
		Pairs: make(map[Expression]Expression),
	}

	for !parser.PeekTokenIs(RBRACE) {
		parser.NT()
		key := parser.parseExpression(LOWEST)

		if !parser.ExpectPeek(COLON) {
			return nil
		}

		parser.NT()
		value := parser.parseExpression(LOWEST)
		hash.Pairs[key] = value

		if !parser.PeekTokenIs(RBRACE) && !parser.ExpectPeek(COMMA) {
			return nil
		}
	}

	if !parser.ExpectPeek(RBRACE) {
		return nil
	}

	return hash
}

func (parser *Parser) parseMacroLiteral() Expression {
	tok := parser.CurrentToken

	if !parser.ExpectPeek(LPAREN) {
		return nil
	}

	params := parser.parseFunctionParameters()

	if !parser.ExpectPeek(LBRACE) {
		return nil
	}

	body := parser.parseBlockStatement()

	return &MacroLiteral{
		Token:      tok,
		Parameters: params,
		Body:       body,
	}
}

func (parser *Parser) parseMethodCallExpression(obj Expression) Expression {
	methodcall := &ObjectCallExpression{
		Token:  parser.CurrentToken,
		Object: obj,
	}
	parser.NT()
	name := parser.parseIdent()
	parser.NT()
	methodcall.Call = parser.parseCallExpression(name)
	return methodcall
}

func (parser *Parser) ParseWhileExpression() Expression {
	EXPR := &WhileCondition{Token: parser.CurrentToken}
	parser.NT()
	EXPR.Condition = parser.parseExpression(LOWEST)
	if !parser.ExpectPeek(LBRACE) {
		return nil
	}
	EXPR.Body = parser.parseBlockStatement()
	return EXPR
}
