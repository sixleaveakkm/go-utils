package cutf8

import (
	"bufio"
	"errors"
	"io"
	"unicode"
	"unicode/utf8"
)

// Valid checks if the byte slice is valid UTF-8.
// Standard utf8 validate only checks the byte range,
// this function additionally checks if the byte slice is printable (expect following char:
// '\n', '\r', '\t', 0x03 (ETX), 0x1a (EOF)).
func Valid(bs []byte) bool {
	if !utf8.Valid(bs) {
		return false
	}
	s := string(bs)
	return ValidString(s)
}

// validCommonRune checks if the rune is printable or one of the following characters:
// Standard utf8 validate only checks the byte range,
// this function additionally checks if the byte slice is printable (expect following char:
// '\n', '\r', '\t', 0x03 (ETX), 0x1a (EOF)).
func validCommonRune(r rune) bool {
	return unicode.IsPrint(r) || r == '\n' || r == '\r' || r == '\t' || r == 0x03 || r == 0x1a
}

// ValidRune checks if the rune is utf8, along with if it is printable or one of the following characters:
// Standard utf8 validate only checks the byte range,
// this function additionally checks if the byte slice is printable (expect following char:
// '\n', '\r', '\t', 0x03 (ETX), 0x1a (EOF)).
func ValidRune(r rune) bool {
	return utf8.ValidRune(r) || validCommonRune(r)
}

// ValidString checks if the string is utf8, along with if it is printable or one of the following characters:
// Standard utf8 validate only checks the byte range,
// this function additionally checks if the byte slice is printable (expect following char:
// '\n', '\r', '\t', 0x03 (ETX), 0x1a (EOF)).
func ValidString(s string) bool {
	if !utf8.ValidString(s) {
		return false
	}
	for _, r := range s {
		if !validCommonRune(r) {
			return false
		}
	}
	return true
}

func ValidReader(reader io.Reader) bool {
	buf := bufio.NewReader(reader)
	for {
		r, _, err := buf.ReadRune()
		if errors.Is(err, io.EOF) {
			return true
		}
		if err != nil {
			return false
		}
		if r == unicode.ReplacementChar {
			return false
		}
		if !validCommonRune(r) {
			return false
		}
	}
}
