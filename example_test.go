package utils_test

import (
	"fmt"
	"github.com/sixleaveakkm/go-utils"
)

func ExampleNull() {
	nullMap := make(map[string]utils.Null)

	nullMap["foo"] = utils.None
	_, fooOK := nullMap["foo"]
	_, barOK := nullMap["bar"]
	fmt.Println(fooOK)
	fmt.Println(barOK)
	// Output: true
	// false
}
