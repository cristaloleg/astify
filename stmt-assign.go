package astify

import "go/ast"

// Assign ...
type Assign struct {
	statement
	dst        []*Value
	src        []*Value
	isCreation bool
}

func newAssign(dst, src []ast.Expr) *Assign {
	a := &Assign{
		dst: make([]*Value, 0, len(dst)),
		src: make([]*Value, 0, len(src)),
	}
	for _, d := range dst {
		_ = d
		// a.dst = append(a.dst, newValue(nil))
	}
	for _, s := range src {
		_ = s
		// a.src = append(a.src, newValue(nil))
	}
	return a
}

// IsSingle ...
func (a *Assign) IsSingle() bool {
	return len(a.dst) == 1
}
