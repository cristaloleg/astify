package astify

import (
	"go/parser"
	"go/token"
	"sort"
)

// ParseFile ...
func ParseFile(path string) (*GoFile, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil, err
	}
	return newFile(path, f), nil
}

// Parse ...
func Parse(path string) (*Astify, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, 0)
	if err != nil {
		return nil, err
	}

	a := &Astify{}
	for name, pkg := range pkgs {
		p := &Pkg{
			name:  name,
			path:  path,
			files: make([]*GoFile, 0, len(pkg.Files)),
		}

		// sort files by name to make file order deterministic
		fnames := make([]string, 0, len(pkg.Files))
		for fname := range pkg.Files {
			fnames = append(fnames, fname)
		}
		sort.Strings(fnames)

		for _, fname := range fnames {
			p.files = append(p.files, newFile(fname, pkg.Files[fname]))
		}

		a.pkgs = append(a.pkgs, p)
	}
	return a, nil
}
