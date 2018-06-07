package astify

import "go/ast"

// Comment ...
type Comment struct {
	declaration
	lines  []string
	parent Node
}

func newComment(*ast.CommentGroup) *Comment {
	c := &Comment{}
	return c
}
