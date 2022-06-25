package errors

import "errors"

var (
	// ErrInfoIsTooShort indicates if info is too short.
	ErrInfoIsTooShort = errors.New("too short")
	// ErrInfoTrailingSpaces indicates if info has trailing spaces.
	ErrInfoTrailingSpaces = errors.New("trailing spaces")
	// ErrDirectoryDoesNotExist indicates if migrations directory doesn't exist.
	ErrDirectoryDoesNotExist = errors.New("doesn't exist")
	// ErrDirectoryIsAFile indicates if migrations directory is a file.
	ErrDirectoryIsAFile = errors.New("is a file")
)
