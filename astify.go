package astify

import (
	"go/ast"
	"path/filepath"
)

// Astify ...
type Astify struct {
	name  string
	path  string
	files []*GoFile
	pkgs  []*Pkg
}

// Walk ...
func (a *Astify) Walk(visiter func(file *GoFile, n Node) error) error {
	for _, p := range a.pkgs {
		for _, f := range p.files {
			for _, n := range f.Nodes() {
				if err := visiter(f, n); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Pkg ...
type Pkg struct {
	name  string
	path  string
	files []*GoFile
}

// Walk ...
func (p *Pkg) Walk(visiter func(file *GoFile, n Node) error) error {
	for _, f := range p.files {
		for _, n := range f.Nodes() {
			if err := visiter(f, n); err != nil {
				return err
			}
		}
	}
	return nil
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

func newFile(path string, file *ast.File) *GoFile {
	name, _ := filepath.Abs(path)
	f := &GoFile{
		name: name,
		path: path,
	}
	for _, d := range file.Decls {
		f.nodes = append(f.nodes, conv2Node(d))
	}
	return f
}

func conv2Node(decl ast.Decl) Node {
	return nil
}

// Walk ...
func (f *GoFile) Walk(walk func(Node) error) error {
	for _, d := range f.nodes {
		if err := walk(d); err != nil {
			return err
		}
	}
	return nil
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
