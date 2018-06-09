package astify

import "go/ast"

// Defer ...
type Defer struct {
	statement
}

func newDefer(*ast.CallExpr) *Defer {
	g := &Defer{}
	return g
}
