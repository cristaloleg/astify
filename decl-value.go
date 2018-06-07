package astify

import "go/ast"

// Value ...
type Value struct {
	declaration
	comment *Comment
	name    string
	typ     *GoType
	isConst bool
}

func newValue(spec *ast.ValueSpec) *Value {
	v := &Value{
		name: spec.Names[0].Name,
	}
	return v
}
