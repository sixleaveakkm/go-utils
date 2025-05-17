package str

// FirstN returns the first n characters of s.
// Using rune to support multi-byte characters.
// With range check, out of range free.
func FirstN(s string, n int) string {
	arr := []rune(s)
	if len(arr) < n {
		return s
	}
	return string(arr[:n])
}
