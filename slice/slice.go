package slice // Package slice provides

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	. "github.com/sixleaveakkm/go-utils"
)

func First[T any](param []T) T {
	if len(param) > 0 {
		return param[0]
	}
	var t T
	return t
}

func FirstOr[T any](param []T, def T) T {
	if len(param) > 0 {
		return param[0]
	}
	return def
}

// FirstMaxN is an alternative of arr[:N] to get first maximum N element of given slice, to avoid slice's behaviour on arr[:N] will panic if N is out or range.
// If length of slice is less than N, will get the original slice.
func FirstMaxN[T any](ss []T, maxN int) []T {
	size := maxN
	if len(ss) < maxN {
		size = len(ss)
	}
	return ss[:size]
}

// After is an alternative of arr[N:] to get all element after N, to avoid slice's behaviour on arr[N:] will panic if N is out or range.
func After[T any](ss []T, n int) []T {
	if n >= len(ss) {
		return []T{}
	}
	return ss[n:]
}

// Range is an alternative of arr[from:to] to get all element between from and to, to avoid slice's behaviour on arr[from:to] will panic if from or to is out or range.
func Range[T any](ss []T, from int, to int) []T {
	if from >= len(ss) {
		return []T{}
	}
	if to > len(ss) {
		to = len(ss)
	}
	return ss[from:to]
}

// Union joins two slices into one slice removing duplicate elements (both duplicate inside one slice and duplicate cross two slices).
// For slices which elements are not comparable, use UnionFunc instead.
func Union[T comparable](a []T, b []T) []T {
	m := make(map[T]Null)
	for _, t := range a {
		m[t] = None
	}
	for _, t := range b {
		m[t] = None
	}
	var r []T
	for e := range m {
		r = append(r, e)
	}
	return r
}

// UnionFunc join two slice to one slice, with a function parameter to identify the key of each element.
// Order will not be guaranteed.
// For slices which elements are comparable, use Union instead.
func UnionFunc[T any, K comparable](a []T, b []T, keyFn func(element T) K) []T {
	m := make(map[K]T)
	for _, t := range a {
		k := keyFn(t)
		m[k] = t
	}
	for _, t := range b {
		k := keyFn(t)
		m[k] = t
	}

	var r []T
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// Exclude takes two slices and returns a slice with elements in a but not in b.
// For slices which elements are not comparable, use ExcludeFunc instead.
// Same elements in a will be kept.
func Exclude[T comparable](a []T, b []T) []T {
	m := make(map[T]Null)
	for _, t := range b {
		m[t] = None
	}
	var r []T
	for _, t := range a {
		_, ok := m[t]
		if ok {
			continue
		}
		r = append(r, t)
	}
	return r
}

func ExcludeFunc[T any, K comparable](a []T, b []T, keyFn func(element T) K) []T {
	m := make(map[K]Null)
	for _, t := range b {
		k := keyFn(t)
		m[k] = None
	}
	var r []T
	for _, t := range a {
		k := keyFn(t)
		_, ok := m[k]
		if ok {
			continue
		}
		r = append(r, t)
	}
	return r
}

func ContainsItem[T comparable](s []T, o T) bool {
	for i := range s {
		if s[i] == o {
			return true
		}
	}
	return false
}

// Set takes a slice and returns a slice remove following appeared duplicate elements.
func Set[T comparable](a []T) []T {
	m := make(map[T]Null)
	var r []T
	for _, t := range a {
		_, ok := m[t]
		if ok {
			continue
		}
		m[t] = None
		r = append(r, t)
	}
	return r
}

// FromSingleton returns a slice of given element.
func FromSingleton[T any](e T) []T {
	return []T{e}
}

// EqualsNoOrder compares two slices without order.
// It is a wrapper of github.com/google/go-cmp/cmp.Equal with cmpopts.SortSlices(less).
func EqualsNoOrder[T any](a, b []T, less func(x, y T) bool) bool {
	return cmp.Equal(a, b, cmpopts.SortSlices(less))
}
