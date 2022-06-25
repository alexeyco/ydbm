package generator

import "github.com/spf13/afero"

// Validator describes validator interface.
type Validator interface {
	Validate(afero.Fs, string, string) error
}
