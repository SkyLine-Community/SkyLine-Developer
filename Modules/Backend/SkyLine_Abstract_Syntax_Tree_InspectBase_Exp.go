package SkyLine_Backend

import (
	"fmt"
	"strconv"
)

func (b *Builtin) Inspect() string        { return "builtin function" }                                // Methodize Inspect  | Inspect built in function call
func (s *String) Inspect() string         { return s.Value }                                           //  Methodize Inspect | Inspect string data type or value
func (q *Quote) Inspect() string          { return fmt.Sprintf("%s(%s)", QuoteType, q.Node.String()) } //  Methodize Inspect | Inspect Quote type and value
func (e *Error) Inspect() string          { return e.Message }                                         //  Methodize Inspect | Inspect error type
func (rv *ReturnValue) Inspect() string   { return rv.Value.Inspect() }                                //  Methodize Inspect | Inspect return values
func (n *Nil) Inspect() string            { return "" }                                                //  Methodize Inspect | Inspect empty/NULL/0x00 value
func (f *Float) Inspect() string          { return strconv.FormatFloat(f.Value, 'f', -1, 64) }         //  Methodize Inspect | Inspect Float types
func (i *Integer) Inspect() string        { return strconv.FormatInt(i.Value, 10) }                    //  Methodize Inspect | Inspect integer types
func (ui *Unsigned) Inspect() string      { return strconv.FormatUint(ui.Value, 10) }                  //  Methodize Inspect | Inspect Unsigned types
func (b *Boolean_Object) Inspect() string { return strconv.FormatBool(b.Value) }                       //  Methodize Inspect | Inspect boolean values
