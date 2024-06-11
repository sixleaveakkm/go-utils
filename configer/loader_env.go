package configer

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type EnvLoader[T any] struct {
	PrepareFunc []func() error
}

func NewEnvLoader[T any]() *EnvLoader[T] {
	return &EnvLoader[T]{}
}

func (e *EnvLoader[T]) AddPrepareFunc(f func() error) *EnvLoader[T] {
	e.PrepareFunc = append(e.PrepareFunc, f)
	return e
}

func (e *EnvLoader[T]) Parse(c *T) error {
	type tuple struct {
		v   *T
		err error
	}

	res := make(chan tuple, 1)
	go func() {
		var v *T
		for _, f := range e.PrepareFunc {
			err := f()
			if err != nil {
				res <- tuple{v: nil, err: err}
			}
		}

		err := env.Parse(v)
		res <- tuple{v: v, err: err}
	}()

	r := <-res
	if r.err != nil {
		return r.err
	}

	if r.v == nil {
		return fmt.Errorf("unexpected parse error: result is nil")
	}
	*c = *r.v
	return nil
}
