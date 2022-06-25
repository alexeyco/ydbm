package rules

import (
	"fmt"
	"strings"

	"github.com/spf13/afero"

	"github.com/alexeyco/ydbm/internal/generator/validator/errors"
)

// InfoIsTrimmed checks if info doesn't have trailing spaces.
func InfoIsTrimmed(_ afero.Fs, info, _ string) error {
	trimmed := strings.TrimSpace(info)
	if trimmed == info {
		return nil
	}

	return fmt.Errorf("migration info %q shouldn't have %w", info, errors.ErrInfoTrailingSpaces)
}
