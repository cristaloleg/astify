package astify

import "go/ast"

// Param ...
type Param struct {
	name string
	typ  *GoType
}

func newParam(name string, typ *GoType) *Param {
	p := &Param{
		name: name,
		typ:  typ,
	}
	return p
}

// Name ...
func (p *Param) Name() string {
	return p.name
}

// Type ...
func (p *Param) Type() *GoType {
	return p.typ
}

// IsExported ...
func (p *Param) IsExported() bool {
	return ast.IsExported(p.name)
}

// IsNamed ...
func (p *Param) IsNamed() bool {
	return len(p.name) != 0
}

// IsSkipped ...
func (p *Param) IsSkipped() bool {
	return p.name == "_"
}
