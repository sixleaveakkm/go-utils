package strutil

import "strings"

func (m *Manipulator) TrimSpace() *Manipulator {
	s := strings.TrimSpace(string(m.data))
	m.data = []byte(s)
	return m
}

// RemoveSpace removes all spaces from the string.
// It does not remove other whitespace characters like tabs or newlines, judging by `strings.IsSpace`.
func (m *Manipulator) RemoveSpace() *Manipulator {
	s := strings.Join(strings.Fields(string(m.data)), "")
	m.data = []byte(s)
	return m
}
