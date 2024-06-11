package utils_test

import (
	"fmt"
	. "github.com/sixleaveakkm/go-utils"
	"github.com/sixleaveakkm/go-utils/errz"
	"github.com/stretchr/testify/assert"
	"testing"
)

const str = "foo"

func TestResult(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		fn1 := func() Result[string] {
			return OK(str)
		}
		r1 := fn1()
		assert.True(t, r1.IsOK())
		assert.False(t, r1.IsError())
		assert.Equal(t, str, r1.Unwrap())
	})

	t.Run("err", func(t *testing.T) {
		fn2 := func() Result[string] {
			return Err[string](fmt.Errorf("failed"))
		}
		r2 := fn2()
		assert.False(t, r2.IsOK())
		assert.True(t, r2.IsError())
		assert.Equal(t, "failed", r2.UnwrapError().Error())
	})

	t.Run("from and multiErr", func(t *testing.T) {
		fn3 := func() (string, error) {
			return str, nil
		}
		fn4 := func() (string, error) {
			return "", fmt.Errorf("failed")
		}
		r3 := ResultFrom(fn3())
		r4 := ResultFrom(fn4())
		assert.True(t, r3.IsOK())
		assert.True(t, r4.IsError())

		errs := errz.New()
		assert.NoError(t, errs.Err())
		v3 := r3.Unwraps(errs)
		v3_1 := r3.Unwraps(errs)
		_ = r4.Unwraps(errs)
		_ = r4.Unwraps(errs)

		assert.Equal(t, str, v3)
		assert.Equal(t, str, v3_1)
		assert.Equal(t, 2, errs.Len())
		// response:
		//     2 errors:
		//     <tab>failed
		//     <tab>failed
		fmt.Println(errs.Err())

	})

	t.Run("solo err", func(t *testing.T) {
		// cannot infer T
		// r5 := Err(fmt.Errorf("failed"))
		r5 := Err[string](fmt.Errorf("failed"))

		assert.Panics(t, func() {
			r5.Unwrap()
		})

		v5 := r5.UnwrapUncheck()
		assert.Zero(t, v5)

		bar := "bar"
		assert.Equal(t, bar, r5.UnwrapOr(bar))
	})
}
