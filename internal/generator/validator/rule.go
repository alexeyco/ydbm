package validator

import "github.com/spf13/afero"

// Rule describes validation rule.
type Rule func(afero.Fs, string, string) error
