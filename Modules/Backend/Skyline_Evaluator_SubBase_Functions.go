package SkyLine_Backend

func OBJECT_CONV_TO_NATIVE_BOOL(o Object) bool {
	if r, ok := o.(*ReturnValue); ok {
		o = r.Value
	}
	switch obj := o.(type) {
	case *Boolean_Object:
		return obj.Value
	case *String:
		return obj.Value != ""
	case *Nil:
		return false
	case *Integer:
		if obj.Value == 0 {
			return false
		}
		return true
	case *Float:
		if obj.Value == 0.0 {
			return false
		}
		return true
	case *Array:
		if len(obj.Elements) == 0 {
			return false
		}
		return true
	case *Hash:
		if len(obj.Pairs) == 0 {
			return false
		}
		return true
	default:
		return true
	}
}
