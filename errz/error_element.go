package errz

import "fmt"

type errorElement struct {
	Index int
	Msg   string
	Err   error
}

// Error error interface implementation.
func (ee errorElement) Error() string {
	return fmt.Sprintf("outside index: %d, with msg: '%s' has error: %+v", ee.Index, ee.Msg, ee.Err)
}

// Unwrap error's Is, As support.
func (ee errorElement) Unwrap() error {
	return ee.Err
}
