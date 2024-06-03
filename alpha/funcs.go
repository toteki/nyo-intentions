package alpha

// Trying to write an add function which sets c=a+b, in such a way that nil values are possible.

type Int *int64

func NewInt(n int64) Int {
	return Int(&n) // TODO: does this mutate weirdly?
}

func Add(a Int, b Int, c Int) {
	sum := int64(0)
	if a != nil {
		sum = sum + *a
	}
	if b != nil {
		sum = sum + *b
	}
	c = &sum // does this actually mutate c?
	// *c = sum // or would this?
}
