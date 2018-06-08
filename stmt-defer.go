package astify

// Defer ...
type Defer struct {
	statement
}

func newDefer(...interface{}) *Defer {
	g := &Defer{}
	return g
}
