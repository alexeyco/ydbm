package rules

import (
	"fmt"
	"unicode/utf8"

	"github.com/spf13/afero"

	"github.com/alexeyco/ydbm/internal/generator/validator/errors"
)

const infoMinLen = 5

// InfoLength checks migration info length.
func InfoLength(_ afero.Fs, info, _ string) error {
	l := utf8.RuneCountInString(info)
	if l >= infoMinLen {
		return nil
	}

	return fmt.Errorf("migration info is %w; should contain at least %d characters", errors.ErrInfoIsTooShort, infoMinLen)
}
