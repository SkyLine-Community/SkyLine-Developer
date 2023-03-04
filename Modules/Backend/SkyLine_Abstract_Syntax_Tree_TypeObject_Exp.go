package SkyLine_Backend

func (i *Integer) Type_Object() Type_Object        { return IntegerType }     // Object Interface   | Integer type object
func (f *Float) Type_Object() Type_Object          { return FloatType }       // Object Interface   | Float type object
func (b *Boolean_Object) Type_Object() Type_Object { return BooleanType }     // Object Interface   | Boolean type object
func (n *Nil) Type_Object() Type_Object            { return NilType }         // Object Interface   | Null type object
func (rv *ReturnValue) Type_Object() Type_Object   { return ReturnValueType } // Object Interface   | return value type object
func (e *Error) Type_Object() Type_Object          { return ErrorType }       // Object Interface   | Error type object
func (f *Function) Type_Object() Type_Object       { return FunctionType }    // Object Interface   | Function type object
func (s *String) Type_Object() Type_Object         { return StringType }      // Object Interface   | String type object
func (b *Builtin) Type_Object() Type_Object        { return BuiltinType }     // Object Interface   | Built in function type object
func (*Array) Type_Object() Type_Object            { return ArrayType }       // Object Interface   | Array type object
func (*Hash) Type_Object() Type_Object             { return HashType }        // Object Interface   | Hash type object
func (q *Quote) Type_Object() Type_Object          { return QuoteType }       // Object Interface   | Quote type object
func (m *Macro) Type_Object() Type_Object          { return MacroType }       // Object Interface   | Macro type object
func (ui *Unsigned) Type_Object() Type_Object      { return UnsignedType }    // Object Interface   | Unsigned type object
