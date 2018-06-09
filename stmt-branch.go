package astify

import (
	"go/ast"
	"go/token"
)

// Branch ...
type Branch struct {
	statement
	isBreak       bool
	isGoto        bool
	isContinue    bool
	isFallthrough bool
}

func newBranch(stmt *ast.Ident, tok token.Token) *Branch {
	b := &Branch{
		isBreak:       tok == token.BREAK,
		isGoto:        tok == token.GOTO,
		isContinue:    tok == token.CONTINUE,
		isFallthrough: tok == token.FALLTHROUGH,
	}
	return b
}
