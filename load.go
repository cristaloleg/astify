package astify

import (
	"fmt"
	"go/parser"
	"log"

	"golang.org/x/tools/go/loader"
)

// Load ...
func Load(importPaths []string) (*Astify, error) {

	conf := loader.Config{
		ParserMode: parser.ParseComments,
		// TypeChecker: types.Config{
		// 	Sizes: sizes,
		// },
	}

	if _, err := conf.FromArgs(importPaths, true); err != nil {
		log.Fatalf("resolve packages: %v", err)
	}
	prog, err := conf.Load()
	if err != nil {
		log.Fatalf("load program: %v", err)
	}

	a := &Astify{}

	for _, pkgPath := range importPaths {

		pkgInfo := prog.Imported[pkgPath]
		if pkgInfo == nil || !pkgInfo.TransitivelyErrorFree {
			log.Fatalf("%s package is not properly loaded", pkgPath)
		}

		p := &Pkg{
			name:  pkgPath, // name,
			path:  pkgPath, //path,
			files: make([]*GoFile, 0, len(pkgInfo.Files)),
		}

		a.typesInfo = &pkgInfo.Info
		fmt.Printf("%v\n", a.typesInfo)
		for _, f := range pkgInfo.Files {
			p.files = append(p.files, newFile(f.Name.Name, f))
		}

		a.pkgs = append(a.pkgs, p)
	}

	return a, nil
}
