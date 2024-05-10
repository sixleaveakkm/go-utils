package ptr

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDefaultOr(t *testing.T) {
	t.Run("DefaultOr for bool", func(t *testing.T) {
		tr := true
		fa := false
		pt := &tr
		pf := &fa

		require.True(t, DefaultOr(pt))
		require.False(t, DefaultOr(pf))
	})
}
