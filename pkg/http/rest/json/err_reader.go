package json

import "errors"

// ErrReader is a mock error that implements io.Reader interface and always returns an error.
type ErrReader int

// Read will return an error.
func (ErrReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("Read failed")
}
