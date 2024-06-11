package range_expression

import (
	"fmt"
	"github.com/sixleaveakkm/go-utils/slice"
	"sort"
	"strconv"
	"strings"
)

var InvalidFormatError = fmt.Errorf("invalid format")

func Parse(expression string) ([]int, error) {
	var activePos []int
	cells := strings.Split(expression, ",")
	for _, cell := range cells {
		if len(cell) == 0 {
			continue
		}
		rngP := strings.Split(cell, "-")
		if len(rngP) > 2 {
			return nil, InvalidFormatError
		}
		start, err := strconv.Atoi(rngP[0])
		if err != nil {
			return nil, InvalidFormatError
		}
		end := start
		if len(rngP) == 2 {
			end, err = strconv.Atoi(rngP[1])
			if err != nil {
				return nil, InvalidFormatError
			}
		}
		if start > end {
			return nil, InvalidFormatError
		}
		for i := start; i <= end; i++ {
			activePos = append(activePos, i)
		}
	}
	activePos = slice.Set(activePos)
	sort.Ints(activePos)
	return activePos, nil
}
