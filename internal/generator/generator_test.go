package generator_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/ydbm/internal/generator"
	"github.com/alexeyco/ydbm/internal/generator/context"
)

func TestGenerator_To(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	const directory = "path/to/directory"

	expectedError := errors.New("error")

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()
		validatorMock := NewMockValidator(ctrl)
		actionMock := NewMockAction(ctrl)

		ctx := context.Context{
			Fs:        fsMock,
			Directory: directory,
			New: context.Migration{
				Version: 1,
				Info:    "foo bar",
				Struct:  "migration0001",
			},
		}

		validatorMock.EXPECT().Validate(fsMock, ctx.New.Info, directory).Return(nil)
		actionMock.EXPECT().Generate(ctx).Return(nil)

		err := generator.Generate(
			ctx.New.Info,
			generator.WithFs(fsMock),
			generator.WithValidator(validatorMock),
			generator.WithActions(actionMock),
		).To(directory)

		assert.NoError(t, err)

		fi, err := fsMock.Stat(directory)

		assert.NoError(t, err)
		assert.True(t, fi.IsDir())
	})

	t.Run("ActionError", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()
		validatorMock := NewMockValidator(ctrl)
		actionMock := NewMockAction(ctrl)

		ctx := context.Context{
			Fs:        fsMock,
			Directory: directory,
			New: context.Migration{
				Version: 1,
				Info:    "foo bar",
				Struct:  "migration0001",
			},
		}

		validatorMock.EXPECT().Validate(fsMock, ctx.New.Info, directory).Return(nil)
		actionMock.EXPECT().Generate(ctx).Return(expectedError)

		err := generator.Generate(
			ctx.New.Info,
			generator.WithFs(fsMock),
			generator.WithValidator(validatorMock),
			generator.WithActions(actionMock),
		).To(directory)

		assert.ErrorIs(t, err, expectedError)
	})

	t.Run("ValidatorError", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()
		validatorMock := NewMockValidator(ctrl)
		actionMock := NewMockAction(ctrl)

		ctx := context.Context{
			Fs:        fsMock,
			Directory: directory,
			New: context.Migration{
				Version: 1,
				Info:    "foo bar",
				Struct:  "migration0001",
			},
		}

		validatorMock.EXPECT().Validate(fsMock, ctx.New.Info, directory).Return(expectedError)

		err := generator.Generate(
			ctx.New.Info,
			generator.WithFs(fsMock),
			generator.WithValidator(validatorMock),
			generator.WithActions(actionMock),
		).To(directory)

		assert.ErrorIs(t, err, expectedError)
	})
}
