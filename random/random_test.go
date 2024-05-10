package random

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestStringAny(t *testing.T) {
	t.Run("utf8 charset", func(t *testing.T) {
		s, err := StringAny("あいうえお", 10)
		require.NoError(t, err)
		r := []rune(s)
		require.Equal(t, len(r), 10)

		expectedChar := false
		for _, it := range r {
			switch it {
			case 'あ', 'い', 'う', 'え', 'お':
				continue
			default:
				expectedChar = true
				break
			}
		}
		require.False(t, expectedChar)
	})

	t.Run("num charset", func(t *testing.T) {
		s, err := String(Numeric, 5)
		require.NoError(t, err)
		_, err = strconv.Atoi(s)
		require.NoError(t, err)
	})
}
