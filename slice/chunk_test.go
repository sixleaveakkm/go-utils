package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	sliceSize := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 97, 98, 99, 100, 101,
	}
	for _, size := range sliceSize {
		t.Run(fmt.Sprintf("array size %d", size), func(t *testing.T) {
			var s []int
			for i := 0; i < size; i++ {
				s = append(s, i)
			}
			res, err := Chunk(s, 10)
			assert.NoError(t, err)
			assert.Len(t, res, (size+9)/10)
			for i := 0; i < len(res); i++ {
				if i == len(res)-1 {
					if size%10 == 0 {
						assert.Len(t, res[i], 10)
					} else {
						assert.Len(t, res[i], size%10)
					}
				} else {
					assert.Len(t, res[i], 10)
				}
			}
			fmt.Println(res)

		})
	}
}
