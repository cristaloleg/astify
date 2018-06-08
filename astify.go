package astify

import (
	"go/ast"
	"path/filepath"
	"regexp"
	"strings"
)

// Astify ...
type Astify struct {
	name string
	path string
	pkgs []*Pkg
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

// Name returns a name of the package.
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
	switch decl := decl.(type) {
	case *ast.FuncDecl:
		return newFunction(decl)

	case *ast.GenDecl:
		for _, spec := range decl.Specs {
			switch spec := spec.(type) {
			case *ast.ImportSpec:
				return newImport(spec)

			case *ast.ValueSpec:
				return newValue(spec)

			case *ast.TypeSpec:
				switch spec.Type.(type) {
				case *ast.StructType:
					return newStruct(spec)

				case *ast.InterfaceType:
					return newInterface(spec)

				case *ast.FuncType:
					// return newFuncType(spec)
				case *ast.MapType:
					// return newMapType(spec)
				case *ast.ChanType:
					// return newChanType(spec)
				case *ast.ArrayType:
					// return newArrayType(spec)
				default:
					return nil
				}

			default:
				println("woah nil")
				return nil
			}
		}
		println("woah nil x2")
		return nil

	default:
		println("woah nil x3")
		return nil
	}
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

// Name returns a name of the Go file.
func (f *GoFile) Name() string {
	return f.name
}

// Path returns a path to the Go file.
func (f *GoFile) Path() string {
	return f.path
}

// Size return a size of the Go file in bytes.
func (f *GoFile) Size() int {
	return f.size
}

// Package ...
func (f *GoFile) Package() *Pkg {
	return f.pkg
}

// Nodes returns a list of nodes inside a Go file.
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

// WalkFlag ...
type WalkFlag int64

const (
	// AllFlag ...
	AllFlag WalkFlag = ^(0)
	// CodeFlag ...
	CodeFlag WalkFlag = iota
	// TestFlag ...
	TestFlag
	// VendorFlag ...
	VendorFlag
	// LinuxFlag ...
	LinuxFlag
	// WindowsFlag ...
	WindowsFlag
	// DarwinFlag ...
	DarwinFlag
	// ...
)

// IsFile ...
func IsFile(file *GoFile, flag WalkFlag) bool {
	switch flag {
	case AllFlag:
		return true

	case CodeFlag:
		return !strings.HasSuffix(file.name, "_test.go")

	case TestFlag:
		return strings.HasSuffix(file.name, "_test.go")

	case VendorFlag:
		return strings.Contains(file.name, "/vendor/")

	case LinuxFlag:
		return strings.HasSuffix(file.name, "_linux.go")

	case WindowsFlag:
		return strings.HasSuffix(file.name, "_windows.go")

	case DarwinFlag:
		return strings.HasSuffix(file.name, "_darwin.go")

	default:
		// log
		return false
	}
}

// IsFilePathMatches ...
func IsFilePathMatches(file *GoFile, re *regexp.Regexp) bool {
	return re.MatchString(file.Path())
}

// IsFileNameMatches ...
func IsFileNameMatches(file *GoFile, re *regexp.Regexp) bool {
	return re.MatchString(file.Name())
}
