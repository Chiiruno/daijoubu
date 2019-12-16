// Package util contains various general utility functions.
package util

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"strings"
)

// WrappedError wraps error types to create compound error chains.
type WrappedError struct {
	Text  string
	Inner error
}

// WrapError wraps error types to create compound error chains.
func WrapError(text string, err error) error {
	return WrappedError{
		Text:  text,
		Inner: err,
	}
}

func (e WrappedError) Error() string {
	if e.Inner != nil {
		e.Text = fmt.Sprintf("%s: %s", e.Text, e.Inner.Error())
	}

	return e.Text
}

// Waterfall executes a slice of functions until the first error returns.
// This error, if any, is returned to the caller.
func Waterfall(fns ...func() error) (err error) {
	for _, fn := range fns {
		if err = fn(); err != nil {
			break
		}
	}

	return
}

// Parallel executes functions in parallel. The first error is returned, if any.
func Parallel(fns ...func() error) error {
	ch := make(chan error)

	for _, fn := range fns {
		go func(fn func() error) {
			ch <- fn()
		}(fn)
	}

	for range fns {
		if err := <-ch; err != nil {
			return err
		}
	}

	return nil
}

// HashBuffer computes a base64 MD5 hash from a buffer.
func HashBuffer(buf []byte) string {
	hash := md5.Sum(buf)
	return base64.RawStdEncoding.EncodeToString(hash[:])
}

// ConcatStrings efficiently concatenates strings with only one extra allocation.
func ConcatStrings(s ...string) string {
	l := 0

	for _, s := range s {
		l += len(s)
	}

	b := make([]byte, 0, l)

	for _, s := range s {
		b = append(b, s...)
	}

	return string(b)
}

// CloneBytes creates a copy of b.
func CloneBytes(b []byte) []byte {
	cp := make([]byte, len(b))
	copy(cp, b)
	return cp
}

// SplitPunctuation splits off one byte of leading and trailing punctuation,
// if any, and returns the 3 split parts. If there is no edge punctuation, the
// respective byte = 0.
func SplitPunctuation(word []byte) (leading byte, mid []byte, trailing byte) {
	mid = word

	// Split leading
	if len(mid) < 2 {
		return
	}

	if isPunctuation(mid[0]) {
		leading = mid[0]
		mid = mid[1:]
	}

	// Split trailing
	l := len(mid)

	if l < 2 {
		return
	}

	if isPunctuation(mid[l-1]) {
		trailing = mid[l-1]
		mid = mid[:l-1]
	}

	return
}

// Returns, if b is a punctuation symbol.
func isPunctuation(b byte) bool {
	switch b {
	case '!', '"', '\'', '(', ')', ',', '-', '.', ':', ';', '?', '[', ']':
		return true
	default:
		return false
	}
}

// SplitPunctuationString splits off one byte of leading and trailing
// punctuation, if any, and returns the 3 split parts. If there is no edge
// punctuation, the respective byte = 0.
func SplitPunctuationString(word string) (
	leading byte, mid string, trailing byte,
) {
	// Generic copy paste :^)
	// Generic copy paste :^)
	mid = word

	// Split leading
	if len(mid) < 2 {
		return
	}

	if isPunctuation(mid[0]) {
		leading = mid[0]
		mid = mid[1:]
	}

	// Split trailing
	l := len(mid)

	if l < 2 {
		return
	}

	if isPunctuation(mid[l-1]) {
		trailing = mid[l-1]
		mid = mid[:l-1]
	}

	return
}

// GetIP extracts the IP of a request.
func GetIP(r *http.Request) (ip net.IP, err error) {
	var s string
	h := r.Header.Get("X-Forwarded-For")

	if h != "" {
		if i := strings.LastIndexByte(h, ','); i != -1 {
			h = h[i+1:]
		}

		// Header can contain padding spaces
		s = strings.TrimSpace(h)
	}

	if s == "" {
		s = r.RemoteAddr
	}

	if split, _, err := net.SplitHostPort(s); err == nil {
		s = split // Port in address
	} else {
		err = nil
	}

	ip = net.ParseIP(s)

	if ip == nil {
		err = fmt.Errorf("invalid IP: %s", s)
	}

	return
}
