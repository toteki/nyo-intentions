package charlie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// This file should test the behavior of charlie.I in ways that can't be accomplished
// using only exported fields and methods.

func TestMutate(t *testing.T) {
	require := require.New(t)

	/*
		// This panics (intentionally) because undefinedI() or I{} should not be used.

		I1 := undefinedI() // nil
		I1.Set(1)
	*/

	// This group of tests works
	I2 := partiallyDefinedI() // &nil
	I2.Set(2)
	expected2A := int64(2)
	expected2B := &expected2A
	expected2C := &expected2B
	require.Equal(expected2C, I2.i)
	require.Equal(expected2B, *I2.i)
	require.Equal(expected2A, **I2.i)

	// This tests what is actually happening when Set occurs
	i3 := int64(0)
	p3 := &i3
	require.Equal(int64(0), *p3)
	I3 := pointerDefinedI(p3)
	I3.Set(3)
	// Observe that p3 is unaffected - in other words, the pointer inside I points to a new int64,
	// but the original int64 does not have its value overwritten
	require.Equal(int64(0), *p3)
	// Confirm again that Set works
	expected3A := int64(3)
	expected3B := &expected3A
	expected3C := &expected3B
	require.Equal(expected3C, I3.i)
	require.Equal(expected3B, *I3.i)
	require.Equal(expected3A, **I3.i)

	// This test ensures we can set and unset repeatedly
	I4 := fullyDefinedI(4)
	expected4A := int64(4)
	expected4B := &expected4A
	expected4C := &expected4B
	require.Equal(expected4C, I4.i)
	require.Equal(expected4B, *I4.i)
	require.Equal(expected4A, **I4.i)
	I4.Set(5)
	expected5A := int64(5)
	expected5B := &expected5A
	expected5C := &expected5B
	require.Equal(expected5C, I4.i)
	require.Equal(expected5B, *I4.i)
	require.Equal(expected5A, **I4.i)
	I4.Unset()
	require.NotNil(I4.i)
	require.Nil(*I4.i)
	I4.Set(6)
	expected6A := int64(6)
	expected6B := &expected6A
	expected6C := &expected6B
	require.Equal(expected6C, I4.i)
	require.Equal(expected6B, *I4.i)
	require.Equal(expected6A, **I4.i)
}

func TestIs(t *testing.T) {
	require := require.New(t)

	// This undefined I should never appear in practice, but if it does, it still Isnt
	I := undefinedI()
	require.False(I.Is())
	require.True(I.Isnt())

	// Undefined I we intend to use
	I = partiallyDefinedI()
	require.False(I.Is())
	require.True(I.Isnt())

	// Defined I
	I = fullyDefinedI(0)
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
}

// make an I by providing a pointer, for testing.
// This can be used to confirm behavior which may or may not mutate the pointer.
func pointerDefinedI(p *int64) I {
	return I{i: &p}
}

// make a fully defined I without using its constructor, for testing.
func fullyDefinedI(n int64) I {
	p := &n
	return I{i: &p}
}

// make a partially defined I without using its constructor, for testing.
func partiallyDefinedI() I {
	var p *int64    // nil
	return I{i: &p} // i is NOT nil, but it points to nil
}

// make a fully undefined I without using its constructor, for testing.
// We learned in TestMutate that these are useless for value-receiver methods.
func undefinedI() I {
	return I{}
}
