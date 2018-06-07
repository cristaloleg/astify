package astify

// Astify ...
type Astify struct {
	name  string
	path  string
	files []*GoFile
	pkgs  []*Pkg
}

// Pkg ...
type Pkg struct {
	name  string
	path  string
	files []*GoFile
}

// Name ...
func (p *Pkg) Name() string {
	return p.name
}

// GoFile ...
type GoFile struct {
	declaration
	name  string
	path  string
	size  int
	pkg   *Pkg
	nodes []Node
}

// Name ...
func (f *GoFile) Name() string {
	return f.name
}

// Path ...
func (f *GoFile) Path() string {
	return f.path
}

// Size ...
func (f *GoFile) Size() int {
	return f.size
}

// Package ...
func (f *GoFile) Package() *Pkg {
	return f.pkg
}

// Nodes ...
func (f *GoFile) Nodes() []Node {
	return f.nodes
}

// Node ...
type Node interface {
	node()
}

type node struct {
	// pos, end int
}

type (
	declaration node
	statement   node
	expression  node
)

func (d *declaration) node() {}
func (s *statement) node()   {}
func (e *expression) node()  {}
