package charlie_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/toteki/nyo-intentions/charlie"
)

func TestIs(t *testing.T) {
	require := require.New(t)

	// Undefined I we intend to use
	I := charlie.Int()
	require.False(I.Is())
	require.True(I.Isnt())

	// Defined I
	I = charlie.NewInt(0)
	require.True(I.Is())
	require.False(I.Isnt())

	// Make I undefined
	I.Unset()
	require.False(I.Is())
	require.True(I.Isnt())

	// Make I defined
	I.Set(1)
	require.True(I.Is())
	require.False(I.Isnt())

	// Conclusions: This is actually all I wanted in terms of the behavior of the above methods.
	// However, because I{} leads to panics (voluntarily; or the inability to Set otherwise),
	// the I class fails to meet the objective of "empty struct is a (working) undefined".

	// To be more specific, I want that rule to be true in Nyo code, not necessarily the output Go.
	// If Nyo requires initialization (defined or not) and thus cannot produce the empty struct, then we're fine.

	// But if that's acceptable, then equally acceptable is using *int64 directly in the output Go, right?
	// As long as I can write Nyo methods for it, I do not need struct{**int64} in Go. Well, then I can't
	// write methods for it at all in Go, but I can write functions which take what would be the method
	// receiver as the first arg instead.

	// That will be my next step. It is needed for other things I had planned too, such as initializer methods
	// where the receiver must be nil, and just execute-if-non-nil in general.

	// The Go code will use functions only, no methods, and check for nils. The Nyo can allow them as methods.

	// Alternatively, I can continue along the charlie path, just so I can keep the Set and Unset methods
	// (and later, Equals, GT, LTE, et cetera) in the Go output. Methods-only (or functions-only) avoids
	// built-in operations like +,-,== and the like.

	// Either way, looks like I'm trying arithmetic next.
}
