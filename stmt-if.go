package astify

import "go/ast"

// If ...
type If struct {
	statement
}

func newIf(ast.Expr, *ast.BlockStmt, ast.Stmt) *If {
	i := &If{}
	return i
}
