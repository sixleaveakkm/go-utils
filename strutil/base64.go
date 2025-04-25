package strutil

import (
	"encoding/base64"
	"errors"
	"fmt"
)

type Base64Op int

const (
	Base64EncodeStdPadding Base64Op = iota
	Base64EncodeStdNoPad
	Base64EncodeUrlPadding
	Base64EncodeUrlNoPad
)

const (
	Base64Decode Base64Op = iota + 10
	Base64DecodeStdPadding
	Base64DecodeStdNoPad
	Base64DecodeUrlPadding
	Base64DecodeUrlNoPad
)

var ErrInvalidBase64Op = errors.New("invalid base64 operation")

func (m *Manipulator) Base64(op Base64Op) *Manipulator {
	if m.err != nil {
		return m
	}

	var err error
	switch op {
	case Base64EncodeStdPadding:
		m.data = []byte(base64.StdEncoding.EncodeToString(m.data))
	case Base64EncodeStdNoPad:
		m.data = []byte(base64.RawStdEncoding.EncodeToString(m.data))
	case Base64EncodeUrlPadding:
		m.data = []byte(base64.URLEncoding.EncodeToString(m.data))
	case Base64EncodeUrlNoPad:
		m.data = []byte(base64.RawURLEncoding.EncodeToString(m.data))
	case Base64Decode:
		m.data, err = base64.StdEncoding.DecodeString(string(m.data))
		if m.err == nil {
			return m
		}
		m.data, err = base64.RawStdEncoding.DecodeString(string(m.data))
		if m.err == nil {
			return m
		}
		m.data, err = base64.URLEncoding.DecodeString(string(m.data))
		if m.err == nil {
			return m
		}
		m.data, err = base64.RawURLEncoding.DecodeString(string(m.data))
	case Base64DecodeStdPadding:
		m.data, err = base64.StdEncoding.DecodeString(string(m.data))
	case Base64DecodeStdNoPad:
		m.data, err = base64.RawStdEncoding.DecodeString(string(m.data))
	case Base64DecodeUrlPadding:
		m.data, err = base64.URLEncoding.DecodeString(string(m.data))
	case Base64DecodeUrlNoPad:
		m.data, err = base64.RawURLEncoding.DecodeString(string(m.data))
	default:
		m.err = ErrInvalidBase64Op
	}
	if err != nil {
		m.err = fmt.Errorf("base64 operation failed: %w", err)
	}
	return m
}
