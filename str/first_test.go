package str

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstN(t *testing.T) {
	t.Run("alpha", func(t *testing.T) {
		r := FirstN("hello", 3)
		assert.Equal(t, "hel", r)

		r = FirstN("hello", 10)
		assert.Equal(t, "hello", r)
	})

	t.Run("JA", func(t *testing.T) {
		r := FirstN("こんにちわ", 3)
		assert.Equal(t, "こんに", r)

		r = FirstN("日本", 1)
		assert.Equal(t, "日", r)

		r = FirstN("グローバル", 3)
		assert.Equal(t, "グロー", r)
	})

}
