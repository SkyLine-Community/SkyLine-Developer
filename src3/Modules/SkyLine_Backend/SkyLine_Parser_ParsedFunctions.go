package SkyLine

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func (parser *Parser) parseStatement() Statement {
	switch parser.CurrentToken.Token_Type {
	case IMPORT:
		return parser.ParseImportStatement()
	case CARRIER:
		return parser.ParseCarrierStatement()
	case LET:
		return parser.parseLetStatement()
	case RETURN:
		return parser.parseReturnStatement()
	default:
		return parser.parseExpressionStatement()
	}
}

func (parser *Parser) parseLetStatement() *LetStatement {
	stmt := &LetStatement{Token: parser.CurrentToken}
	if !parser.ExpectPeek(IDENT) {
		return nil
	}

	stmt.Name = &Ident{
		Token: parser.CurrentToken,
		Value: parser.CurrentToken.Literal,
	}

	if !parser.ExpectPeek(ASSIGN) {
		return nil
	}

	parser.NextToken()

	stmt.Value = parser.parseExpression(LOWEST)

	for parser.PeekTokenIs(SEMICOLON) {
		parser.NextToken()
	}

	return stmt
}

func (parser *Parser) ParseImportStatement() *Import {
	statement := &Import{Token: parser.CurrentToken}
	if !parser.PeekTokenIs("(") {
		fmt.Println("Missing ( in import statement")
	}
	parser.NextToken()
	statement.ImportValue = parser.parseExpression(LOWEST)
	for parser.PeekTokenIs(SEMICOLON) {
		parser.NextToken()
	}
	currentfilebody := FileCurrent.Get_Body(false)
	smallerfile := "Lexer_Creation_EXEC_MASHED.csc"
	var NewFileBody []string
	var ImportFileBody []string
	FileCurrent.New(statement.ImportValue.String())
	if ok := FileCurrent.VerifyFileExists(false); ok {
		if iscsc := FileCurrent.Verify_CSC(false); iscsc {
			if oktoparse := FileCurrent.Verify_GoodToparse(false); oktoparse {
				// create and store data from current file in
				ImportFileBody = FileCurrent.Get_Body(false)
				if currentfilebody != nil && ImportFileBody != nil {
					for i := 0; i < len(ImportFileBody); i++ {
						NewFileBody = append(NewFileBody, ImportFileBody[i])
					}
					for i := 0; i < len(currentfilebody); i++ {
						NewFileBody = append(NewFileBody, currentfilebody[i])
					}
					remove1 := `import("%s")`
					remove2 := `require("%s")`
					remove3 := `include("%s")`
					remove1 = fmt.Sprintf(remove1, statement.ImportValue.String())
					remove2 = fmt.Sprintf(remove2, statement.ImportValue.String())
					remove3 = fmt.Sprintf(remove3, statement.ImportValue.String())
					newarr := []string{}
					for _, k := range NewFileBody {
						if k != remove1 && k != remove2 && k != remove3 {
							newarr = append(newarr, k)
						}
					}
					// parse and remove file
					FileCurrent.New(smallerfile)
					f, x := os.Create(FileCurrent.Get_Name())
					if x != nil {
						log.Fatal(x)
					}
					defer f.Close()
					for _, char := range newarr {
						if _, x = f.WriteString(char + "\n"); x != nil {
							log.Fatal(x)
						}
					}
					// now execute said file
					_, e := ReadImportedCarrierFile(FileCurrent.Get_Name())
					if e != nil {
						log.Fatal(e)
					}
					FileCurrent.Delete()
					os.Exit(0)
					//FileCurrent.New(smallerfile)
					//_, e := ReadImportedCarrierFile(FileCurrent.Filename)
					//if e != nil {
					//	log.Fatal(e)
					//}
				}
			}
		}
	}
	return statement
}

func (parser *Parser) ParseCarrierStatement() *Carrier {
	statement := &Carrier{Token: parser.CurrentToken}
	if !parser.PeekTokenIs("|") {
		fmt.Println(ErrorSymBolMap[CODE_EXPECT_PEEK_ERROR_DURING_CALL_TO_PEEK](parser.CurrentToken.Literal, "|"))
		return nil
	}
	parser.NextToken()
	statement.CarrierValue = parser.parseExpression(LOWEST)
	for parser.PeekTokenIs(SEMICOLON) {
		parser.NextToken()
	}
	// this is a very VERY minimal linker if a linker at all, this takes the carry statement, parses the file within it and executes the code within that file
	// IT DOES NOT MASH THE SOURCE OF THAT FILE INTO THE MAIN FILE BEING RUN YET!
	FileCurrent.New(statement.CarrierValue.String())
	if ok := FileCurrent.VerifyFileExists(false); ok {
		// Now verify is csc
		if iscsc := FileCurrent.Verify_CSC(false); iscsc {
			// Load and run file
			if oktoparse := FileCurrent.Verify_GoodToparse(false); oktoparse {
				_, x := ReadImportedCarrierFile(statement.CarrierValue.String())
				if x != nil {
					log.Fatal(x)
				}
			}
		} else {
			fmt.Println("Warning: SkyLine Carrier FROM -->  carry|'", FileCurrent.Get_Name(), "|' => File is not a .csc file or Cyber Security Core file, can not run this file")
		}

	} else {
		if FileCurrent.IsDir {
			fmt.Println("Warning: SkyLine Carrier (File IS NOT A FILE, this was a directirt ) CARRY does not include directories yet....ERR: ", FileCurrent.Get_Name())
		} else {
			fmt.Println("Warning: SkyLine Carrier (File does not exist ->  ", FileCurrent.Get_Name(), " ) Please ensure the filename is correct")
		}
	}
	return statement
}

func (parser *Parser) parseReturnStatement() *ReturnStatement {
	stmt := &ReturnStatement{
		Token: parser.CurrentToken,
	}
	parser.NextToken()

	stmt.ReturnValue = parser.parseExpression(LOWEST)

	for parser.PeekTokenIs(SEMICOLON) {
		parser.NextToken()
	}

	return stmt
}

func (parser *Parser) parseExpressionStatement() *ExpressionStatement {
	stmt := &ExpressionStatement{
		Token:      parser.CurrentToken,
		Expression: parser.parseExpression(LOWEST),
	}

	if parser.PeekTokenIs(SEMICOLON) {
		parser.NextToken()
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

		parser.NextToken()

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

//TODO: Remake this error system for 90%
func (parser *Parser) parseIntegerLiteral() Expression {
	lit := &IntegerLiteral{Token: parser.CurrentToken}

	val, err := strconv.ParseInt(parser.CurrentToken.Literal, 0, 64)
	if err != nil {
		fmt.Println(ErrorSymBolMap[CODE_PARSE_INT_ERROR](parser.CurrentToken.Literal))
	}

	lit.Value = val
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

	parser.NextToken()

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

	parser.NextToken()

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
	parser.NextToken()

	expr := parser.parseExpression(LOWEST)

	if !parser.ExpectPeek(LINE) {
		return nil
	}

	return expr
}

func (parser *Parser) parseGroupedExpression() Expression {
	parser.NextToken()

	expr := parser.parseExpression(LOWEST)

	if !parser.ExpectPeek(RPAREN) {
		return nil
	}

	return expr
}

func (parser *Parser) parseWhileExpression() Expression {
	EXPR := &WhileCondition{Token: parser.CurrentToken}
	parser.NextToken()
	EXPR.Condition = parser.parseExpression(LOWEST)
	if !parser.ExpectPeek(LBRACE) {
		return nil
	}
	EXPR.Body = parser.parseBlockStatement()
	return EXPR
}

func (parser *Parser) parseIfExpression() Expression {
	expr := &ConditionalExpression{Token: parser.CurrentToken}

	parser.NextToken()
	expr.Condition = parser.parseExpression(LOWEST)

	if !parser.ExpectPeek(LBRACE) {
		return nil
	}

	expr.Consequence = parser.parseBlockStatement()

	if parser.PeekTokenIs(ELSE) {
		parser.NextToken()

		if !parser.ExpectPeek(LBRACE) {
			return nil
		}

		expr.Alternative = parser.parseBlockStatement()
	}

	return expr
}

func (parser *Parser) parseBlockStatement() *BlockStatement {
	block := &BlockStatement{
		Token:      parser.CurrentToken,
		Statements: []Statement{},
	}

	parser.NextToken()
	for !parser.CurrentTokenIs(RBRACE) && !parser.CurrentTokenIs(EOF) {
		stmt := parser.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}

		parser.NextToken()
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
		parser.NextToken()
		return idents
	}

	parser.NextToken()

	ident := &Ident{
		Token: parser.CurrentToken,
		Value: parser.CurrentToken.Literal,
	}
	idents = append(idents, ident)

	for parser.PeekTokenIs(COMMA) || parser.PeekTokenIs(COLON) {
		parser.NextToken()
		parser.NextToken()
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
		parser.NextToken()
		return list
	}

	parser.NextToken()
	list = append(list, parser.parseExpression(LOWEST))

	for parser.PeekTokenIs(COMMA) {
		parser.NextToken()
		parser.NextToken()
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

	parser.NextToken()
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
		parser.NextToken()
		key := parser.parseExpression(LOWEST)

		if !parser.ExpectPeek(COLON) {
			return nil
		}

		parser.NextToken()
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
