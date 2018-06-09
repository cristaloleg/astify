package astify

import "go/ast"

// Body ...
type Body struct {
	statement
	stmts []Node
}

func newBody(body *ast.BlockStmt) Node {
	b := &Body{
		stmts: make([]Node, 0, len(body.List)),
	}
	for _, s := range body.List {
		b.stmts = append(b.stmts, conv2Stmt(s))
	}
	return b
}

func conv2Stmt(stmt ast.Stmt) Node {
	switch stmt := stmt.(type) {
	case *ast.BadStmt:
		// log

	case *ast.DeclStmt:
	case *ast.EmptyStmt:
	case *ast.LabeledStmt:
	case *ast.ExprStmt:
	case *ast.SendStmt:
	case *ast.CaseClause:
	case *ast.SwitchStmt:
	case *ast.TypeSwitchStmt:
	case *ast.CommClause:
	case *ast.ForStmt:
	case *ast.RangeStmt:

	case *ast.SelectStmt:
		return newSelect(stmt.Body)

	case *ast.IncDecStmt:
		return newIncDec(stmt)

	case *ast.BranchStmt:
		return newBranch(stmt.Label, stmt.Tok)

	case *ast.AssignStmt:
		return newAssign(stmt.Lhs, stmt.Rhs)

	case *ast.GoStmt:
		return newGoroutine(stmt.Call)

	case *ast.DeferStmt:
		return newDefer(stmt.Call)

	case *ast.BlockStmt:
		return conv2Stmt(stmt)

	case *ast.IfStmt:
		return newIf(stmt.Cond, stmt.Body, stmt.Else)

	case *ast.ReturnStmt:
		for _, r := range stmt.Results {
			switch expr := r.(type) {
			case *ast.BinaryExpr:
				_ = expr
			}
		}
	default:
		return nil
	}
	return nil
}
