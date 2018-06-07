package astify

import "go/ast"

// GoType ...
type GoType struct {
	name        string
	isPointer   bool
	isBuiltin   bool
	isPrimitive bool
	isReference bool
	isAlias     bool
}

func newGoType(interface{}) *GoType {
	gt := &GoType{}
	return gt
}

// Name ...
func (gt *GoType) Name() string {
	return gt.name
}

// IsEqual ...
func (gt *GoType) IsEqual(other *Type) bool {
	return false
}

// IsPointerOf ...
func (gt *GoType) IsPointerOf(other *Type) bool {
	return false
}

// IsPointerTo ...
func (gt *GoType) IsPointerTo(other *Type) bool {
	return false
}

// IsExported ...
func (gt *GoType) IsExported() bool {
	return ast.IsExported(gt.name)
}

// IsPointer ...
func (gt *GoType) IsPointer() bool {
	return gt.isPointer
}

// IsBuiltin ...
func (gt *GoType) IsBuiltin() bool {
	return gt.isBuiltin
}

// IsPrimitive ...
func (gt *GoType) IsPrimitive() bool {
	return gt.isPrimitive
}

// IsAlias ...
func (gt *GoType) IsAlias() bool {
	return gt.isAlias
}
