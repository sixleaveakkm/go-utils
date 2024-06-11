package errz_test

import (
	"errors"
	"fmt"
	"github.com/sixleaveakkm/go-utils/errz"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorElement(t *testing.T) {
	var err0 error
	errs := errz.New()

	errs.AppendWithIndex(err0, 0)
	err1 := io.ErrUnexpectedEOF
	errs.AppendWithIndex(err1, 1)
	assert.Equal(t, 1, errs.Len())
	e := errs.First()
	assert.True(t, errors.Is(e, io.ErrUnexpectedEOF))

}

func TestErrors(t *testing.T) {
	errs := errz.New()
	assert.Equal(t, 0, errs.Len())
	foo, bar := "foo", "bar"
	fn1 := func() (string, error) {
		return foo + bar, nil
	}
	fn2 := func(x int, y string) (string, error) {
		return y, fmt.Errorf("cannot concat int with string") // for example
	}
	res := errz.Or(errs, fn1)
	assert.Equal(t, "foobar", res)
	assert.Equal(t, 0, errs.Len())
	assert.NoError(t, errs.Err())

	res2 := errz.Or(errs, func() (string, error) {
		return fn2(1, foo)
	})
	assert.Equal(t, foo, res2)
	assert.Equal(t, 1, errs.Len())

	res3 := errz.Or(errs, func() (string, error) {
		return fn2(1, foo)
	})
	assert.Equal(t, "", res3)
	assert.Equal(t, 1, errs.Len())

	res4 := errz.Do(errs, func() (string, error) {
		return fn2(1, foo)
	})
	assert.Equal(t, "foo", res4)
	assert.Equal(t, 2, errs.Len())
}
