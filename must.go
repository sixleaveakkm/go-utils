package utils

// Must takes the (value, error) pair and return the value if error is nil.
// Panic if error is not nil.
//
// Example:
//
//	func Foo() (int, error) {
//	    return 1, nil
//	}
//
// i := utils.Must(Foo())
func Must[T any](res T, err error) T {
	if err != nil {
		panic(err)
	}
	return res
}
