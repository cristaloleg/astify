package astify

// func AsPackage(n Node) *Package
// 	pack, ok := n.(*Package)
// if !ok{
// 		panic("not a package")
// }
// 	return pack
// }

// AsImport ...
func AsImport(n Node) *Import {
	imp, ok := n.(*Import)
	if !ok {
		panic("not an Import")
	}
	return imp
}

// AsComment ...
func AsComment(n Node) *Comment {
	comm, ok := n.(*Comment)
	if !ok {
		panic("not a Comment")
	}
	return comm
}

// AsFunction ...
func AsFunction(n Node) *Function {
	fn, ok := n.(*Function)
	if !ok {
		panic("not a Function")
	}
	return fn
}

// AsStruct ...
func AsStruct(n Node) *Struct {
	strct, ok := n.(*Struct)
	if !ok {
		panic("not a Struct")
	}
	return strct
}

// AsInterface ...
func AsInterface(n Node) *Interface {
	iface, ok := n.(*Interface)
	if !ok {
		panic("not an Interface")
	}
	return iface
}

// AsType ...
func AsType(n Node) *Type {
	typ, ok := n.(*Type)
	if !ok {
		panic("not a Type")
	}
	return typ
}

// AsValue ...
func AsValue(n Node) *Value {
	value, ok := n.(*Value)
	if !ok {
		panic("not a Value")
	}
	return value
}
