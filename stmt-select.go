package astify

import "go/ast"

// Select ...
type Select struct {
	statement
}

func newSelect(body *ast.BlockStmt) *Select {
	s := &Select{}
	return s
}
