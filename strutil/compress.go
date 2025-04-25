package strutil

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"fmt"
	"io"
)

type CompressOp int

const (
	// CompressOpFlate use the default compression.
	CompressOpFlate CompressOp = iota
	CompressOpGzip
)

func (m *Manipulator) Compress(op CompressOp) *Manipulator {
	if m.err != nil {
		return m
	}

	var buf bytes.Buffer
	var w io.WriteCloser
	var err error

	switch op {
	case CompressOpFlate:
		w, err = flate.NewWriter(&buf, flate.DefaultCompression)
	case CompressOpGzip:
		w = gzip.NewWriter(&buf)
	default:
		m.err = fmt.Errorf("invalid compression operation: %d", op)
		return m
	}
	if err != nil {
		m.err = fmt.Errorf("failed to create compressor: %w", err)
		return m
	}
	if _, err = w.Write(m.data); err != nil {
		m.err = fmt.Errorf("failed to write data: %w", err)
		return m
	}
	if err = w.Close(); err != nil {
		m.err = fmt.Errorf("failed to close compressor: %w", err)
		return m
	}
	m.data = buf.Bytes()
	return m
}

func (m *Manipulator) Decompress(op CompressOp) *Manipulator {
	if m.err != nil {
		return m
	}

	var buf bytes.Buffer
	var r io.ReadCloser
	var err error

	switch op {
	case CompressOpFlate:
		r = flate.NewReader(bytes.NewReader(m.data))
	case CompressOpGzip:
		r, err = gzip.NewReader(bytes.NewReader(m.data))
	default:
		m.err = fmt.Errorf("invalid decompression operation: %d", op)
		return m
	}
	if err != nil {
		m.err = fmt.Errorf("failed to create decompressor: %w", err)
		return m
	}
	if _, err = io.Copy(&buf, r); err != nil {
		m.err = fmt.Errorf("failed to read data: %w", err)
		return m
	}
	if err = r.Close(); err != nil {
		m.err = fmt.Errorf("failed to close decompressor: %w", err)
		return m
	}
	m.data = buf.Bytes()
	return m
}
