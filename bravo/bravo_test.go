package bravo_test

import (
	"testing"

	"github.com/toteki/nyo-intentions/bravo"
)

func TestMutate(t *testing.T) {
	// Non-nil define still requires a literal number
	a := bravo.NewInt(3)

	// Mutliple means of defining as nil. Redundant?
	b := bravo.Int()
	var c bravo.I

	if a.Value() == nil {
		t.Log("a was nil")
		t.FailNow()
	}

	// Tese never reaches here
	b.Unset()
	c.Unset()

	// Conclusions:
	// NewInt(3) failed because it called func (i I) Set(n) {i.i = &n} which modified the copy of i sent into
	// the function, but not in the calling code.

	// Next Steps:
	// For charlie.I, I will be forced to use func (i *I) Set(n int64) and defend against i == nil some other way.
	// Also, do fields within I need to be pointers anymore? Probably yes, so more nil checks for field access.
}
