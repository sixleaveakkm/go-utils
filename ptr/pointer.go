package ptr

import (
	"fmt"
	"reflect"
)

func String(s string) *string {
	return &s
}

func Stringf(f string, a ...any) *string {
	s := fmt.Sprintf(f, a...)
	return &s
}

func Bool(b bool) *bool {
	return &b
}

func Int64(i int64) *int64 {
	return &i
}

func Int(i int) *int {
	return &i
}

// DefaultOr get value of a pointer. If pointer is nil, zero value of this type will be return.
func DefaultOr[T any](i *T) T {
	var v T
	if i == nil {
		return v
	}
	return *i
}

// DefaultOrElse get value of a pointer. If pointer is nil, or value will be return.
func DefaultOrElse[T any](i *T, or T) T {
	if i == nil {
		return or
	}
	return *i
}

// Ptr return a pointer of the value
func Ptr[T any](i T) *T {
	return &i
}

// PtrOrNilForZero return a pointer of given not-zero value
// If given value is a zero value (reflect IsZero), nil will be return
func PtrOrNilForZero[T any](i T) *T {
	if reflect.ValueOf(i).IsZero() {
		return nil
	}
	return &i
}
