package files

import "io"

// Storage defines the behaviour of file operations
// Implementation is like local storage, cloude storage etc.
type Storage interface {
	Save(path string, file io.Reader) error
}
