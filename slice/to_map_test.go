package slice_test

import (
	"fmt"
	"github.com/sixleaveakkm/go-utils/slice"
)

func ExampleToMap() {
	arr := []string{"1", "2", "3"}
	// convert slice to map with key as index, value os value.
	m1 := slice.ToMap(arr, func(idx int, v string) int {
		return idx
	})
	fmt.Println(m1)

	type Complex struct {
		ID   int
		Name string
	}

	arr2 := []Complex{
		{ID: 1, Name: "foo"},
		{ID: 2, Name: "bar"},
	}
	// convert slice to map with some value in slice item.
	m2 := slice.ToMap(arr2, func(_ int, v Complex) int {
		return v.ID
	})
	fmt.Println(m2)

	// output:
	// map[0:1 1:2 2:3]
	// map[1:{1 foo} 2:{2 bar}]
}

func ExampleToMapWithDuplicateCheck() {
	type Complex struct {
		ID   int
		Name string
	}

	arr := []Complex{
		{ID: 1, Name: "foo"},
		{ID: 2, Name: "bar"},
		{ID: 2, Name: "foo bar"},
	}

	getKey := func(_ int, v Complex) int {
		return v.ID
	}

	onDupKeepExist := func(key int, existValue Complex, newValue Complex) (Complex, error) {
		return existValue, nil
	}

	onDupDoErr := func(key int, existValue Complex, newValue Complex) (Complex, error) {
		return Complex{}, fmt.Errorf("duplicate on key %d, existValue: %v, newValue: %v", key, existValue, newValue)
	}
	m1, err := slice.ToMapWithDuplicateCheck(arr, getKey, onDupKeepExist)
	fmt.Println(m1, ", ", err)

	m2, err := slice.ToMapWithDuplicateCheck(arr, getKey, onDupDoErr)
	fmt.Println(m2, ", ", err)

	// output:
	// map[1:{1 foo} 2:{2 bar}] ,  <nil>
	// map[] ,  duplicate on key 2, existValue: {2 bar}, newValue: {2 foo bar}
}

func ExampleToSetMap() {
	arr := []string{"1", "2", "3", "10", "1", "2"}

	setMap := slice.ToSetMap(arr)
	fmt.Println(setMap)

	for v := range setMap {
		// do something with unique elements in arr
		_ = v
	}

}
