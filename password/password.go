package password

import (
	"crypto/rand"
	"github.com/sixleaveakkm/go-utils/random"
	"math/big"
	"strings"
	"unicode/utf8"
)

// Policy is password policy contains
// - Validate tests if given password fits this policy
// - GenerateTempPassword generates a new temporary password fits this policy.
type Policy struct {
	MinimumLength           int
	RequireNumber           bool
	RequireUppercase        bool
	RequireLowercase        bool
	RequireSpecialCharacter bool
	SpecialCharacters       string
}

func requireNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func requireUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func requireLower(r rune) bool {
	return r >= 'a' && r <= 'a'
}

// Validate tests if given password fits this policy
func (p Policy) Validate(password string) bool {
	password = strings.TrimSpace(password)

	if len(password) < p.MinimumLength {
		return false
	}

	if p.RequireSpecialCharacter && len(p.SpecialCharacters) > 0 && !strings.ContainsAny(password, p.SpecialCharacters) {
		return false
	}

	if p.RequireNumber && !strings.ContainsFunc(password, requireNumber) {
		return false
	}
	if p.RequireUppercase && !strings.ContainsFunc(password, requireUpper) {
		return false
	}
	if p.RequireLowercase && !strings.ContainsFunc(password, requireLower) {
		return false
	}
	return true
}

// GenerateTempPassword generates a new temporary password fits this policy.
// This function just fit the policy, doesn't have full randomness.
func (p Policy) GenerateTempPassword() (string, error) {
	s := random.MustString(random.LowerAlphabet, p.MinimumLength)
	if p.RequireNumber {
		s += random.MustString(random.Numeric, 1)
	}
	if p.RequireUppercase {
		s += random.MustString(random.UpperAlphabet, 1)
	}
	if p.RequireSpecialCharacter {
		c, err := random.StringAny(p.SpecialCharacters, 1)
		if err != nil {
			return "", err
		}
		s += c
	}

	runeSlice := []rune(s)
	n := utf8.RuneCountInString(s)

	for i := n - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return "", err
		}
		k := int(j.Int64())
		runeSlice[i], runeSlice[k] = runeSlice[k], runeSlice[i]
	}

	return string(runeSlice), nil
}
