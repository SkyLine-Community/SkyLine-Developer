package SkyLine_Backend

func (ls *LetStatement) SN()        {} // Statement Node     | Allow condition
func (rs *ReturnStatement) SN()     {} // Statement Node     | Allow Return
func (es *ExpressionStatement) SN() {} // Statement Node     | Allow Expression
func (Const *Constant) SN()         {} // Statement Node     | Allow Constants
