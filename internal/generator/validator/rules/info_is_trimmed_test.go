package rules_test

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"

	"github.com/alexeyco/ydbm/internal/generator/validator/errors"
	"github.com/alexeyco/ydbm/internal/generator/validator/rules"
)

func TestInfoIsTrimmed(t *testing.T) {
	t.Parallel()

	fsMock := afero.NewMemMapFs()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		err := rules.InfoIsTrimmed(fsMock, "info", "")

		require.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		err := rules.InfoIsTrimmed(fsMock, " info", "")

		require.ErrorIs(t, err, errors.ErrInfoTrailingSpaces)
	})
}
