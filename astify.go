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

// GoFile ...
type GoFile struct {
	declaration
	name  string
	path  string
	size  int
	pkg   *Pkg
	nodes []Node
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
