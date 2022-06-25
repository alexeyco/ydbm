package validator

import (
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/afero"

	"github.com/alexeyco/ydbm/internal/generator/validator/rules"
)

// Validator describes generator validator.
type Validator []Rule

// New returns new Validator.
func New() *Validator {
	return &Validator{
		rules.InfoLength,
		rules.InfoIsTrimmed,
		rules.DirectoryIsCorrect,
	}
}

// Validate validates generator input args.
func (v Validator) Validate(fs afero.Fs, info, directory string) (res error) {
	for _, rule := range v {
		if err := rule(fs, info, directory); err != nil {
			res = multierror.Append(res, err)
		}
	}

	return
}
