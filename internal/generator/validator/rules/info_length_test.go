package rules_test

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/ydbm/internal/generator/validator/errors"
	"github.com/alexeyco/ydbm/internal/generator/validator/rules"
)

func TestInfoLength(t *testing.T) {
	t.Parallel()

	fsMock := afero.NewMemMapFs()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		err := rules.InfoLength(fsMock, "12345", "")

		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		err := rules.InfoLength(fsMock, "1234", "")

		assert.ErrorIs(t, err, errors.ErrInfoIsTooShort)
	})
}
