package astify

// Branch ...
type Branch struct {
	statement
}

func newBranch(...interface{}) *Branch {
	b := &Branch{}
	return b
}
