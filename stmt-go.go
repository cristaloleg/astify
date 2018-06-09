package astify

import "go/ast"

// Goroutine ...
type Goroutine struct {
	statement
}

func newGoroutine(*ast.CallExpr) *Goroutine {
	g := &Goroutine{}
	return g
}
