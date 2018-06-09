package astify

// Param ...
type Param struct {
	name       string
	typ        *GoType
	isEllipsis bool
}

func newParam(name string, typ *GoType) *Param {
	p := &Param{
		name: name,
		typ:  typ,
	}
	return p
}

// Name returns a name of the param.
func (p *Param) Name() string {
	return p.name
}

// Type ...
func (p *Param) Type() *GoType {
	return p.typ
}

// IsNamed ...
func (p *Param) IsNamed() bool {
	return len(p.name) != 0
}

// IsSkipped ...
func (p *Param) IsSkipped() bool {
	return p.name == "_"
}

// IsEllipsis ...
func (p *Param) IsEllipsis() bool {
	return p.isEllipsis
}
