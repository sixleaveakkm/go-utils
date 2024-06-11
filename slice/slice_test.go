package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstOr(t *testing.T) {
	t.Run("param is nil (int)", func(t *testing.T) {
		var s []int
		res := FirstOr(s, 100)
		assert.Equal(t, 100, res)
	})

	t.Run("param is nil (string)", func(t *testing.T) {
		var s []string
		res := FirstOr(s, "foo")
		assert.Equal(t, "foo", res)
	})

	t.Run("param one element (int)", func(t *testing.T) {
		s := []int{1}
		res := FirstOr(s, 100)
		assert.Equal(t, 1, res)
	})
	t.Run("param one element (string)", func(t *testing.T) {
		s := []string{"foo"}
		res := FirstOr(s, "bar")
		assert.Equal(t, "foo", res)
	})

	t.Run("param contains extra elements (int)", func(t *testing.T) {
		s := []int{1, 2, 3}
		res := FirstOr(s, 100)
		assert.Equal(t, 1, res)
	})
	t.Run("param contains extra elements (string)", func(t *testing.T) {
		s := []string{"foo", "bar", "baz"}
		res := FirstOr(s, "qux")
		assert.Equal(t, "foo", res)
	})
}

func TestHeadMaxN(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var s []int
		res := FirstMaxN(s, 100)
		assert.Len(t, res, 0)
	})
	tt := []struct {
		Size   int
		Max    int
		ExpLen int
	}{
		{10, 100, 10},
		{1, 1, 1},
		{10, 0, 0},
		{100, 10, 10},
		{0, 0, 0},
	}
	for _, cs := range tt {
		t.Run(fmt.Sprintf("%d size, max %d", cs.Size, cs.Max), func(t *testing.T) {
			var s []int
			for i := 0; i < cs.Size; i++ {
				s = append(s, i)
			}
			res := FirstMaxN(s, cs.Max)
			assert.Len(t, res, cs.ExpLen)
			for i, it := range res {
				assert.Equal(t, i, it)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	a := []string{"foo", "bar"}
	b := []string{"x", "y", "z"}
	ab := Union(a, b)
	assert.Len(t, ab, 5)
	c := []string{"bar"}
	ac := Union(a, c)
	assert.Len(t, ac, 2)

	var d []string
	ad := Union(a, d)
	assert.Len(t, ad, 2)
	e := []string{
		"e", "e",
	}
	de := Union(d, e)
	assert.Len(t, de, 1)
}

func TestUnionFunc(t *testing.T) {
	type s struct {
		ID int
	}
	a := []s{
		{1}, {2},
	}
	b := []s{
		{2}, {3},
	}
	ab := UnionFunc(a, b, func(element s) int {
		return element.ID
	})
	assert.Len(t, ab, 3)
}
