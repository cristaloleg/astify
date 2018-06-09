package astify

import (
	"go/ast"
	"strings"
)

// Function represents a function declaration.
// Also may represent a method in interface declaration.
type Function struct {
	declaration
	name        string
	comment     *Comment
	receiver    *Param
	params      []*Param
	results     []*Param
	body        Node
	isExternal  bool
	isIfaceDecl bool
}

func newFunction(fnDecl *ast.FuncDecl) *Function {
	fn := &Function{
		name:        fnDecl.Name.Name,
		comment:     newComment(fnDecl.Doc),
		receiver:    makeReceiver(fnDecl.Recv),
		params:      makeParamList(fnDecl.Type.Params),
		results:     makeParamList(fnDecl.Type.Results),
		body:        newBody(fnDecl.Body),
		isExternal:  fnDecl.Body == nil,
		isIfaceDecl: false,
	}
	return fn
}

func newMethod(field *ast.Field) *Function {
	fn := &Function{
		name:        field.Names[0].Name,
		comment:     newComment(field.Doc),
		receiver:    nil,
		params:      makeParamList(field.Type.(*ast.FuncType).Params),
		results:     makeParamList(field.Type.(*ast.FuncType).Results),
		body:        nil,
		isExternal:  true,
		isIfaceDecl: true,
	}
	return fn
}

// Name returns a name of the function.
func (f *Function) Name() string {
	return f.name
}

// Comment returns an associated comment to the function.
func (f *Function) Comment() *Comment {
	return f.comment
}

// Receiver returns a receiver(if any, otherwise a nil) of the function.
func (f *Function) Receiver() *Param {
	return f.receiver
}

// Params returns params of the function.
func (f *Function) Params() []*Param {
	return f.params
}

// Results returns results of the function.
func (f *Function) Results() []*Param {
	return f.results
}

// IsExported returns true if function is exported.
func (f *Function) IsExported() bool {
	return ast.IsExported(f.name)
}

// IsMethod returns true if function is a method of a struct.
func (f *Function) IsMethod() bool {
	return f.receiver != nil
}

// IsExternal returns true if function doesn't have an implemetation.
func (f *Function) IsExternal() bool {
	return f.isExternal
}

// IsTestFunc returns true if it's a testing function.
func (f *Function) IsTestFunc() bool {
	return strings.HasPrefix(f.name, "Test") &&
		len(f.results) == 0 &&
		len(f.params) == 1 &&
		f.params[0].typ.name == "*testing.T"
}

func makeReceiver(recv *ast.FieldList) *Param {
	if recv == nil {
		return nil
	}
	p := &Param{
		name: recv.List[0].Names[0].Name,
		typ:  newGoType(recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name),
	}
	return p
}

func makeParamList(ps *ast.FieldList) []*Param {
	params := []*Param{}
	if ps != nil {
		for _, p := range ps.List {
			if len(p.Names) == 0 {
				params = append(params, newParam("", newGoType(p.Type)))
				continue
			}
			for _, n := range p.Names {
				params = append(params, newParam(n.Name, newGoType(p.Type)))
			}
		}
	}
	return params
}
