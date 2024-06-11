package slice

import "errors"

func Chunk[T any](items []T, size int) (chunks [][]T, err error) {
	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}
	if len(items) == 0 {
		return chunks, nil
	}
	for size < len(items) {
		items, chunks = items[size:], append(chunks, items[0:size:size])
	}
	return append(chunks, items), nil
}
