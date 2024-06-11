package slice

import (
	. "github.com/sixleaveakkm/go-utils"
)

// MapFunc is the function type for ToMap and ToMapWithDuplicateCheck to get key from slice item.
type MapFunc[K comparable, V any] interface {
	~func(idx int, v V) K
}

// MapFuncOnDuplicate is the function type for ToMapWithDuplicateCheck to handle duplicate key.
type MapFuncOnDuplicate[K comparable, V any] interface {
	~func(key K, existValue V, newValue V) (V, error)
}

// ToMap providers a function to convert slice to map, along with a function set instruction how to get key from slice item.
// If there are duplicate key, the last one will be kept. If you want to handle duplicate key, use ToMapWithDuplicateCheck.instead.
func ToMap[K comparable, V any, F MapFunc[K, V]](l []V, getKeyFn F) map[K]V {
	m := make(map[K]V)
	for i, v := range l {
		k := getKeyFn(i, v)
		m[k] = v
	}
	return m
}

// ToMapWithDuplicateCheck providers a function to convert slice to map, along with a function set instruction how to get key from slice item, and a function to handle duplicate key.
// If the onDuplicateFn returns error, it will stop and return immediately with nil map and the error.
func ToMapWithDuplicateCheck[K comparable, V any, F MapFunc[K, V], D MapFuncOnDuplicate[K, V]](l []V, getKeyFn F, onDuplicateFn D) (map[K]V, error) {
	m := make(map[K]V)
	for i, v := range l {
		k := getKeyFn(i, v)
		e, ok := m[k]
		if !ok {
			m[k] = v
			continue
		}

		cv, err := onDuplicateFn(k, e, v)
		if err != nil {
			return nil, err
		}
		m[k] = cv
	}
	return m, nil
}

// ToSetMap takes a slice of comparable type and returns a map with key as slice item and value as None.
func ToSetMap[K comparable](l []K) map[K]Null {
	m := make(map[K]Null)
	for _, v := range l {
		m[v] = None
	}
	return m
}
