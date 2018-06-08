package astify

import "go/ast"

// Assign ...
type Assign struct {
	statement
	dst        Value
	src        Value
	isCreation bool
}

func newAssign(ast.Expr) *Assign {
	a := &Assign{}
	return a
}
