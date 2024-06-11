package errz

import (
	"fmt"
)

type Errors struct {
	errors []error
}

func New() *Errors {
	return &Errors{}
}

// Append stores provided error to Errors if it is not nil.
func (e *Errors) Append(err error) {
	e.errors = append(e.errors, err)
}

// AppendWithIndex stores provided error with given index to Errors if it is not nil.
// This will be useful when you run in a loop and stores all possible error.
func (e *Errors) AppendWithIndex(err error, index int, msg ...string) {
	if err != nil {
		m := ""
		if len(msg) > 0 {
			m = msg[0]
		}
		e.errors = append(e.errors, &errorElement{
			Index: index,
			Msg:   m,
			Err:   err,
		})
	}
}

// Len returns the number of errors stored in Errors.
func (e *Errors) Len() int {
	l := 0
	for _, err := range e.errors {
		if err != nil {
			l++
		}
	}
	return l
}

func (e *Errors) Unwrap() []error {
	var errs []error
	for _, err := range e.errors {
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (e *Errors) Error() string {
	return e.Err().Error()
}

// Err convert Errors to error.
// Note Errors may contain multiple errors, `Is`, `As` check will not work.
// This function is designed to do nil check.
// e.g., `if errs.Err() != nil`
func (e *Errors) Err() error {
	l := 0
	str := ""
	for i, err := range e.errors {
		if err != nil {
			l++
			str += fmt.Sprintf("#%d: %+v\n", i, err)
		}

	}
	if l == 0 {
		return nil
	}
	str = fmt.Sprintf("%d errors:\n", l) + str
	return fmt.Errorf(str)
}

// First takes the first error stored in Errors or nil.
func (e *Errors) First() error {
	for _, err := range e.errors {
		if err != nil {
			return err
		}
	}
	return nil
}

type ErrFunc[T any] interface {
	~func() (T, error)
}

// Do calls func and appends the error to e, return the value.
// The function is limited to no parameter and the return signature is (value, error).
// Another version for return signature of only error is Do0.
func Do[T any, F ErrFunc[T]](e *Errors, fn F) T {
	v, err := fn()
	e.Append(err)

	return v
}

// Do0 calls func and appends the error to e if it is not nil.
// The function is limited to no parameter and the return signature is error.
func Do0(e *Errors, fn func() error) {
	err := fn()
	e.Append(err)
}

// Or calls func only if Errors contains no error, and appends the error to e, return the value.
// If Errors contains any error, it will not run that function and return a zero value of T.
// The function is limited to no parameter and the return signature is (value, error).
// Another version for return signature of only error is Or0.
func Or[T any, F ErrFunc[T]](e *Errors, fn F) T {
	if e.Err() != nil {
		var zero T
		return zero
	}
	v, err := fn()
	e.Append(err)
	return v
}

// Or0 calls func only if Errors contains no error, and appends the error to e if it is not nil.
// If Errors contains any error, it will not run that function.
// The function is limited to no parameter and the return signature is error.
func Or0(e *Errors, fn func() error) {
	if e.Err() != nil {
		return
	}
	err := fn()
	e.Append(err)
}
