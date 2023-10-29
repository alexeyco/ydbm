package rules_test

import (
	"os"
	"path"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"

	"github.com/alexeyco/ydbm/internal/generator/validator/errors"
	"github.com/alexeyco/ydbm/internal/generator/validator/rules"
)

func TestDirectoryIsCorrect(t *testing.T) {
	t.Parallel()

	const directory = "path/to/directory"

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()

		err := fsMock.MkdirAll(directory, os.ModePerm)
		require.NoError(t, err)

		err = rules.DirectoryIsCorrect(fsMock, "", directory)
		require.NoError(t, err)
	})

	t.Run("ErrDirectoryDoesNotExist", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()

		err := rules.DirectoryIsCorrect(fsMock, "", directory)
		require.ErrorIs(t, err, errors.ErrDirectoryDoesNotExist)
	})

	t.Run("ErrDirectoryIsAFile", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()

		err := fsMock.MkdirAll(path.Dir(directory), os.ModePerm)
		require.NoError(t, err)

		_, err = fsMock.Create(directory)
		require.NoError(t, err)

		err = rules.DirectoryIsCorrect(fsMock, "", directory)
		require.ErrorIs(t, err, errors.ErrDirectoryIsAFile)
	})
}
