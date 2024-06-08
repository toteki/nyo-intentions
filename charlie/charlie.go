package charlie

// try for double pointers. I could win, but at what cost?

// I is either an int64 or undefined
type I struct {
	// Constructs: Since type name (I) and constructor name (Int) must be different in Go,
	// do we find a way to make them the same in Nyo? maybe "// nyotype:Int"

	i **int64

	// A variable of type I cannot be nil, but its field i can be nil OR point to nil.
	// Since I cannot be nil, we can call methods on I (not *I) safely.
	// If we define methods as func (i I) then we can mutate both the value and the pointer.
}

// Int returns a new, undefined I
func Int() I {
	// i.i is not nil, but points to nil
	var p *int64
	return I{&p}
	// Constructs: Do we want Type() to always be a func which returns undefined of a given type?
}

// Set I to a given int64 value.
func (i I) Set(n int64) {
	// I.i must not be nil.
	// creates a pointer to the passed value, and sets i.i to point at that pointer
	if i.i != nil {
		// i.i points to either nil or a *int64
		p := &n
		*i.i = p // * allows mutate through pass-by-value
	} else {
		panic("panic for visibility: nil i.i in I.Set")
	}
	// Constructs: shall we require all Nyo types have this method?
	// If so, how? (required type varies, so maybe generics if allowed in interfaces)
}

// Unset I back to undefined
func (i I) Unset() {
	// I.i must not be nil.
	// the previous pointer will probably be garbage collected
	if i.i != nil {
		*i.i = nil // * allows mutate through pass-by-value
	} else {
		panic("panic for visibility: nil i.i in I.Unset")
	}
	// Constructs: shall we require all Nyo types have this method?
}

// NewInt returns a new I set to a given int64 value
func NewInt(n int64) I {
	i := Int()
	i.Set(n)
	return i
	// Constructs: Do we overload Type(...) to use this as well?
}

// Is returns true if I is defined
func (i I) Is() bool {
	// check for two layers of non-nil
	// test: make sure this pattern can't panic
	return i.i != nil && *i.i != nil
	// Constructs: shall we require all Nyo types have this method?
}

// Isnt returns true if I is undefined
func (i I) Isnt() bool {
	// check for two layers of nil
	// test: make sure this pattern can't panic
	return i.i == nil || *i.i == nil
	// Constructs: is it necessary to have this on the Go side?
}
