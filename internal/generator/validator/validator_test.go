package validator_test

import (
	"errors"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"

	"github.com/alexeyco/ydbm/internal/generator/validator"
)

func rule(err error) validator.Rule {
	return func(_ afero.Fs, _, _ string) error {
		return err
	}
}

func TestValidator_Validate(t *testing.T) {
	t.Parallel()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		v := validator.Validator{
			rule(nil),
		}

		err := v.Validate(afero.NewMemMapFs(), "", "")

		require.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		expectedError := errors.New("error")

		v := validator.Validator{
			rule(expectedError),
		}

		err := v.Validate(afero.NewMemMapFs(), "", "")

		require.ErrorIs(t, err, expectedError)
	})
}
