package astify

// func IsPackage(n Node) bool {
// 	_, ok := n.(*Package)
// 	return ok
// }

// IsImport ...
func IsImport(n Node) bool {
	_, ok := n.(*Import)
	return ok
}

// IsComment ...
func IsComment(n Node) bool {
	_, ok := n.(*Comment)
	return ok
}

// IsFunction ...
func IsFunction(n Node) bool {
	_, ok := n.(*Function)
	return ok
}

// IsStruct ...
func IsStruct(n Node) bool {
	_, ok := n.(*Struct)
	return ok
}

// IsInterface ...
func IsInterface(n Node) bool {
	_, ok := n.(*Interface)
	return ok
}

// IsType ...
func IsType(n Node) bool {
	_, ok := n.(*Type)
	return ok
}

// IsValue ...
func IsValue(n Node) bool {
	_, ok := n.(*Value)
	return ok
}
