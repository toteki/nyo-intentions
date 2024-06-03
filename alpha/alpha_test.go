package alpha_test

import (
	"testing"

	"github.com/toteki/nyo-intentions/alpha"
)

func TestMutate(t *testing.T) {
	var a, b, c alpha.Int
	a = alpha.NewInt(1)

	alpha.Add(a, b, c)

	/*
		if c == nil {
			t.Log("c was nil")
			t.FailNow()
		}
	*/

	// Conclusions:
	// We weren't mutating c by setting "c = &sum" inside Add(c)
	// Also, initializing values without being able to set them equal to numbers (a = 1) is too verbose.

	// Next steps: Maybe fix both of those problems.
	// Have a function bravo.Int(n) which returns a struct. What that struct contains should allow mutation
	// by functions and/or by methods.
}
