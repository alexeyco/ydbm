package templatex

import "errors"

var (
	// ErrTemplate describes error with template compilation.
	ErrTemplate = errors.New("template error")
	// ErrFormatting describes error with formatting.
	ErrFormatting = errors.New("formatting error")
)
