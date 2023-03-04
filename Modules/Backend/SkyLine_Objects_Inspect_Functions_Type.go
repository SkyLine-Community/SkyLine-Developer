package SkyLine_Backend

import (
	"bytes"
	"fmt"
	"strings"
)

func (f *Function) Inspect() string {
	var Out bytes.Buffer
	params := make([]string, 0, len(f.Parameters))
	for _, parser := range f.Parameters {
		params = append(params, parser.String())
	}
	Out.WriteString(fmt.Sprint(f.Type_Object()) + "(")
	Out.WriteString(strings.Join(params, ", "))
	Out.WriteString(") {")
	Out.WriteString(f.Body.String())
	Out.WriteString("}")
	return Out.String()
}

func (a *Array) Inspect() string {
	if a == nil {
		return ""
	}
	elements := make([]string, 0, len(a.Elements))
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}
	var Out bytes.Buffer
	Out.WriteString("")
	Out.WriteString("[")
	Out.WriteString(strings.Join(elements, ", "))
	Out.WriteString("]")
	return Out.String()
}

func (h *Hash) Inspect() string {
	if h == nil {
		return ""
	}
	pairs := make([]string, 0, len(h.Pairs))
	for _, pair := range h.Pairs {
		pairs = append(pairs, pair.Key.Inspect()+": "+pair.Value.Inspect())
	}
	var Out bytes.Buffer
	Out.WriteString("{")
	Out.WriteString(strings.Join(pairs, ", "))
	Out.WriteString("}")
	return Out.String()
}

func (m *Macro) Inspect() string {
	var Out bytes.Buffer
	params := make([]string, 0, len(m.Parameters))
	for _, parser := range m.Parameters {
		params = append(params, parser.String())
	}
	Out.WriteString("macro(")
	Out.WriteString(strings.Join(params, ", "))
	Out.WriteString(") {\n")
	Out.WriteString(m.Body.String())
	Out.WriteString("\n}")
	return Out.String()
}
