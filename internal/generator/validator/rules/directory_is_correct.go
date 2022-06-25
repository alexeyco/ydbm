package rules

import (
	"fmt"
	"os"

	"github.com/spf13/afero"

	"github.com/alexeyco/ydbm/internal/generator/validator/errors"
)

// DirectoryIsCorrect checks if migrations directory is correct.
func DirectoryIsCorrect(fs afero.Fs, _, directory string) error {
	fi, err := fs.Stat(directory)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("migrations directory %q %w", directory, errors.ErrDirectoryDoesNotExist)
		}

		return err
	}

	if fi.IsDir() {
		return nil
	}

	return fmt.Errorf("migrations directory %q %w", directory, errors.ErrDirectoryIsAFile)
}
