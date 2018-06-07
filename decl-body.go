package astify

import "go/ast"

// Body ...
type Body struct {
	statement
	stmts []Node
}

func newBody(body *ast.BlockStmt) Node {
	b := &Body{}
	// for _, s := range body.List {
	// 	b.stmts = append(b.stmts, conv2Stmt(s))
	// }
	return b
}

func conv2Stmt(stmt ast.Stmt) Node {
	s := &statement{}
	return s
}
