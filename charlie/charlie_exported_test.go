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
}
