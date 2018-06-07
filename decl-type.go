package astify

import "go/ast"

// Type ...
type Type struct {
	declaration
	name string
	// comment     *Comment
	typ *GoType
}

func newType(spec *ast.TypeSpec) Node {
	typ := &Type{
		name: spec.Name.Name,
	}
	return typ
}

// Name ...
func (t *Type) Name() string {
	return t.name
}
