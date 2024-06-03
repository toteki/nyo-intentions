package bravo

// I is either an int64 or undefined
type I struct {
	i *int64

	// A variable of type I cannot be nil, but its field i can.
	// Therefore, we should be able to call methods on I (not *I) safely in terms of "cannot access method of nil".

	// However, modifying an I in such methods might be tough: change I.i and the original is probably unaffected.
	// Change *I.i and you can mutate numbers but you can't set i back to undefined in the parent code.. right?
}

// Int returns a new, undefined I
func Int() I {
	return I{}
}

// Set I to a given int64 value
func (i I) Set(n int64) {
	i.i = &n
}

// Unset I back to undefined
func (i I) Unset() {
	i.i = nil
}

// NewInt returns a new I set to a given int64 value
func NewInt(n int64) I {
	i := Int()
	i.Set(n)
	return i
}

// Is returns true if I is defined
func (i I) Is() bool {
	return i.i != nil
}

// Isnt returns true if I is undefined
func (i I) Isnt() bool {
	return i.i == nil
}

// Value exposes the int64 pointer inside I for testing, but we might not want it available in the end
func (i I) Value() *int64 {
	return i.i
}
