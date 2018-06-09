package astify

import (
	"go/ast"
	"go/token"
)

// IncDec ...
type IncDec struct {
	statement
	isInc bool
}

func newIncDec(stmt *ast.IncDecStmt) *IncDec {
	id := &IncDec{
		isInc: stmt.Tok == token.INC,
	}
	return id
}
