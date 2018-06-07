package astify

import "go/ast"

// Struct ...
type Struct struct {
	declaration
	name   string
	fields []*Field
}

func newStruct(spec *ast.TypeSpec) *Struct {
	strct := &Struct{
		name: spec.Name.Name,
	}

	astStruct := spec.Type.(*ast.StructType)
	for _, f := range astStruct.Fields.List {
		fld := newField(f)
		strct.fields = append(strct.fields, fld)
	}
	return strct
}

// Name returns a name of the struct.
func (s *Struct) Name() string {
	return s.name
}

// Fields ...
func (s *Struct) Fields() []*Field {
	return s.fields
}

// Field ...
type Field struct {
	declaration
	name       string
	isEmbedded bool
	tags       []*Tag
}

func newField(f *ast.Field) *Field {
	fld := &Field{
		isEmbedded: f.Names == nil,
	}

	if fld.isEmbedded {
		fld.name = f.Type.(*ast.Ident).Name
	} else {
		fld.name = f.Names[0].Name
	}
	if f.Tag != nil {
		fld.tags = []*Tag{newTag(f.Tag.Value)}
	}
	return fld
}

// Name returns a name of the field.
func (f *Field) Name() string {
	return f.name
}

// IsExported ...
func (f *Field) IsExported() bool {
	return ast.IsExported(f.name)
}

// Tag ...
func (f *Field) Tag() *Tag {
	if f.tags != nil {
		return f.tags[0]
	}
	return newTag("")
}

// IsEmbedded ...
func (f *Field) IsEmbedded() bool {
	return f.isEmbedded
}

// IsSkipped ...
func (f *Field) IsSkipped() bool {
	return f.name == "_"
}

// Tag ...
type Tag struct {
	declaration
	value string
}

func newTag(value string) *Tag {
	tag := &Tag{
		value: value,
	}
	return tag
}

// Value ...
func (t *Tag) Value() string {
	return t.value
}
