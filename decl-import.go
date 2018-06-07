package astify

import "go/ast"

// TODO: some of std imports, is there a better way to do this?
var std = buildSet([]string{"math", "io", "sync", "fmt", "math/big", "net", "net/http"})

// Import ...
type Import struct {
	declaration
	name string
	path string
}

func newImport(spec *ast.ImportSpec) *Import {
	impr := &Import{}
	if spec.Name != nil {
		impr.name = spec.Name.Name
	}
	impr.path = spec.Path.Value
	return impr
}

// Name ...
func (i *Import) Name() string {
	return i.name
}

// Path ....
func (i *Import) Path() string {
	return i.path
}

// IsStd ....
func (i *Import) IsStd() bool {
	_, ok := std[i.path]
	return ok
}

// IsFull ....
func (i *Import) IsFull() bool {
	return i.name == "_"
}

// IsNaked ....
func (i *Import) IsNaked() bool {
	return i.name == "."
}
