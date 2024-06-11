package errz_test

import (
	"fmt"
	"github.com/sixleaveakkm/go-utils/errz"
)

func ExampleNew() {
	errs := errz.New()

	errOnOdd := func(i int) error {
		if i%2 == 0 {
			return nil
		}
		return fmt.Errorf("error on odd number %d", i)
	}

	for i := 0; i < 10; i++ {
		// example of doing something
		err := errOnOdd(i)
		if err != nil {
			errs.Append(err)
			continue
		}

		// do something else
	}

	if errs.Unwrap() != nil {
		fmt.Println(errs.Error())
	}

}
