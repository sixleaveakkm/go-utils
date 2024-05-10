package random

import (
	"crypto/rand"
	"errors"
	"math/big"
)

type RandCharset int

const (
	AlphaNumeric RandCharset = iota
	LowerAlphaNumeric
	UpperAlphaNumeric
	Alphabet
	LowerAlphabet
	UpperAlphabet
	Numeric
)

var randCharset = map[RandCharset]string{
	AlphaNumeric:      "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	LowerAlphaNumeric: "0123456789abcdefghijklmnopqrstuvwxyz",
	UpperAlphaNumeric: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	Alphabet:          "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	UpperAlphabet:     "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	LowerAlphabet:     "abcdefghijklmnopqrstuvwxyz",
	Numeric:           "0123456789",
}

// String generate random string of given charset and length.
// Resulting error if RandCharset is not valid, or from `crypto/rand` package.
func String(set RandCharset, length int) (string, error) {
	charset, ok := randCharset[set]
	if !ok {
		return "", errors.New("undefined RandCharset")
	}

	return StringAny(charset, length)
}

func MustString(set RandCharset, length int) string {
	str, err := String(set, length)
	if err != nil {
		panic(err)
	}
	return str
}

// StringAny generates random string from char
func StringAny(charset string, length int) (string, error) {
	if len(charset) == 0 {
		return "", errors.New("empty charset")
	}
	set := []rune(charset)
	setLen := big.NewInt(int64(len(set)))
	randomString := make([]rune, length)
	for i := range randomString {
		r, err := rand.Int(rand.Reader, setLen)
		if err != nil {
			return "", err
		}

		randomString[i] = set[r.Int64()]
	}

	return string(randomString), nil
}

// Bytes generate random bytes of length
func Bytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	return randomBytes, nil
}
