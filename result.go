package utils

func OK[T any](v T) Result[T] {
	return Result[T]{
		ok: &v,
	}
}

func Err[T any](e error) Result[T] {
	return Result[T]{
		err: e,
	}
}

func ResultFrom[T any](v T, err error) Result[T] {
	return Result[T]{
		ok:  &v,
		err: err,
	}
}

type Result[T any] struct {
	ok  *T
	err error
}

type UnwrapWithErrorHolder interface {
	Append(err error)
}

func (re Result[T]) ToTuple() (T, error) {
	return re.UnwrapUncheck(), re.err
}

func (re Result[T]) Unwrap() T {
	if re.err != nil {
		panic(re.err)
	}
	return re.UnwrapUncheck()
}

func (re Result[T]) UnwrapOr(val T) T {
	if re.err != nil {
		return val
	}
	return re.UnwrapUncheck()
}

func (re Result[T]) UnwrapError() error {
	return re.err
}

func (re Result[T]) Unwraps(h UnwrapWithErrorHolder) T {
	h.Append(re.err)
	return re.UnwrapUncheck()
}

func (re Result[T]) UnwrapUncheck() T {
	if re.ok != nil {
		return *re.ok
	}
	var newV T
	return newV
}

func (re Result[T]) IsOK() bool {
	return re.err == nil
}

func (re Result[T]) IsError() bool {
	return re.err != nil
}
