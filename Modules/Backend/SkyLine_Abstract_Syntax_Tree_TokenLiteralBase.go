package SkyLine_Backend

func (ls *LetStatement) TokenLiteral() string                 { return ls.Token.Literal }         // Token Literal 	   | Returns literal of an allow statement
func (i *Ident) TokenLiteral() string                         { return i.Token.Literal }          // Token Literal 	   | Returns literal of an identifier
func (rs *ReturnStatement) TokenLiteral() string              { return rs.Token.Literal }         // Token Literal 	   | Returns literal of a return statement
func (ui *UnsignedLiteral) TokenLiteral() string              { return ui.Token.Literal }         // Token Literal 	   | Returns literal of a unsigned integer
func (es *ExpressionStatement) TokenLiteral() string          { return es.Token.Literal }         // Token Literal 	   | Returns literal of expression statement
func (il *IntegerLiteral) TokenLiteral() string               { return il.Token.Literal }         // Token Literal 	   | Returns literal of integer
func (fl *FloatLiteral) TokenLiteral() string                 { return fl.Token.Literal }         // Token Literal 	   | Returns literal of a float
func (pe *PrefixExpression) TokenLiteral() string             { return pe.Token.Literal }         // Token Literal 	   | Returns literal of PrefixExpression
func (ie *InfixExpression) TokenLiteral() string              { return ie.Token.Literal }         // Token Literal 	   | Returns literal of InfixExpression
func (b *Boolean_AST) TokenLiteral() string                   { return b.Token.Literal }          // Token Literal 	   | Returns literal of Boolean statement
func (ie *ConditionalExpression) TokenLiteral() string        { return ie.Token.Literal }         // Token Literal 	   | Returns literal of Conditionals
func (bs *BlockStatement) TokenLiteral() string               { return bs.Token.Literal }         // Token Literal 	   | Returns literal of code block statements
func (fl *FunctionLiteral) TokenLiteral() string              { return fl.Token.Literal }         // Token Literal 	   | Returns literal of function
func (ce *CallExpression) TokenLiteral() string               { return ce.Token.Literal }         // Token Literal 	   | Returns literal of CallExpression
func (ml *MacroLiteral) TokenLiteral() string                 { return ml.Token.Literal }         // Token Literal 	   | Returns literal of Macro
func (Switch *Switch) TokenLiteral() string                   { return Switch.Token.Literal }     // Token Literal 	   | Returns literal of switch
func (Case *Case) TokenLiteral() string                       { return Case.Token.Literal }       // Token Literal 	   | Returns literal of Case
func (Assign *AssignmentStatement) TokenLiteral() string      { return Assign.Token.Literal }     // Token Literal        | Returns literal of assignment statement
func (Const *Constant) TokenLiteral() string                  { return Const.Token.Literal }      // Token Literal  	   | Returns literal of a constant
func (ObjectCall *ObjectCallExpression) TokenLiteral() string { return ObjectCall.Token.Literal } // Token Literal        | Returns literal of an object call
func (WhileCond *WhileCondition) TokenLiteral() string        { return WhileCond.Token.Literal }  // Token Literal |   Returns token literal of while conditon's

func (sl *StringLiteral) TokenLiteral() string {
	if sl == nil {
		return ""
	}
	return sl.Token.Literal
}

func (al *ArrayLiteral) TokenLiteral() string {
	if al == nil {
		return ""
	}
	return al.Token.Literal
}

func (ie *IndexExpression) TokenLiteral() string {
	if ie == nil {
		return ""
	}
	return ie.Token.Literal
}

func (hl *HashLiteral) TokenLiteral() string {
	if hl == nil {
		return ""
	}
	return hl.Token.Literal
}

func (prog *Program) TokenLiteral() string {
	if len(prog.Statements) == 0 {
		return ""
	}
	return prog.Statements[0].TokenLiteral()
}
