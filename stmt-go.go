package astify

// Goroutine ...
type Goroutine struct {
	statement
}

func newGoroutine(...interface{}) *Goroutine {
	g := &Goroutine{}
	return g
}
