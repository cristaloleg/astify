package astify

import (
	"go/ast"
	"go/types"
)

// GoType ...
type GoType struct {
	name        string
	typ         types.Type
	isPointer   bool
	isBuiltin   bool
	isBasic     bool
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
func (gt *GoType) IsEqual(other *GoType) bool {
	return types.Identical(gt.typ, other.typ)
}

// IsComparable ...
func (gt *GoType) IsComparable(other *GoType) bool {
	return types.Comparable(gt.typ)
}

// IsNullable ...
func (gt *GoType) IsNullable(other *GoType) bool {
	// return types.hasNil(gt.typ)
	return false
}

// IsPointerOf ...
func (gt *GoType) IsPointerOf(other *Type) bool {
	return types.Identical(gt.typ.Underlying(), other.typ.typ)
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

// IsBasic ...
func (gt *GoType) IsBasic() bool {
	return gt.isBasic
}

// IsAlias ...
func (gt *GoType) IsAlias() bool {
	return gt.isAlias
}
