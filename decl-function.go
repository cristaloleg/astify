package astify

import (
	"go/ast"
	"strings"
)

// Function ...
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

// Name ...
func (f *Function) Name() string {
	return f.name
}

// Comment ...
func (f *Function) Comment() *Comment {
	return f.comment
}

// Receiver ...
func (f *Function) Receiver() *Param {
	return f.receiver
}

// Params ...
func (f *Function) Params() []*Param {
	return f.params
}

// Results ...
func (f *Function) Results() []*Param {
	return f.results
}

// IsExported ...
func (f *Function) IsExported() bool {
	return ast.IsExported(f.name)
}

// IsMethod ...
func (f *Function) IsMethod() bool {
	return f.receiver != nil
}

// IsExternal ...
func (f *Function) IsExternal() bool {
	return f.isExternal
}

// IsTestFunc ...
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
