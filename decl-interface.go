package astify

import "go/ast"

// Interface ...
type Interface struct {
	declaration
	name    string
	methods []*Function
}

func newInterface(spec *ast.TypeSpec) *Interface {
	iface := &Interface{
		name: spec.Name.Name,
	}

	astIface := spec.Type.(*ast.InterfaceType)
	for _, m := range astIface.Methods.List {
		fn := newMethod(m)
		iface.methods = append(iface.methods, fn)
	}
	return iface
}

// Name returns a name of the interface.
func (i *Interface) Name() string {
	return i.name
}

// Methods ...
func (i *Interface) Methods() []*Function {
	return i.methods
}
