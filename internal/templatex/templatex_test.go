package templatex_test

import (
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/ydbm/internal/templatex"
)

var (
	file     = "path/to/file.go"
	template = `package pkg

func {{ .Function }}() bool {
return true
}
`
	data = map[string]any{
		"Function": "SomeFunc",
	}

	raw = `package pkg

func SomeFunc() bool {
return true
}
`
	formatted = `package pkg

func SomeFunc() bool {
	return true
}
`
)

func TestTemplate_Save(t *testing.T) {
	t.Parallel()

	t.Run("WithoutFormatting", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()

		err := templatex.Parse(template).
			Save(file,
				templatex.WithFs(fsMock),
				templatex.WithData(data),
			)
		assert.NoError(t, err)

		f, err := fsMock.Open(file)
		assert.NoError(t, err)

		defer func() { _ = f.Close() }()

		b, err := io.ReadAll(f)
		assert.NoError(t, err)

		assert.Equal(t, raw, string(b))
	})

	t.Run("WitFormatting", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()

		err := templatex.Parse(template).
			Save(file,
				templatex.WithFs(fsMock),
				templatex.WithFormat,
				templatex.WithData(data),
			)
		assert.NoError(t, err)

		f, err := fsMock.Open(file)
		assert.NoError(t, err)

		defer func() { _ = f.Close() }()

		b, err := io.ReadAll(f)
		assert.NoError(t, err)

		assert.Equal(t, formatted, string(b))
	})

	t.Run("ErrFormatting", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()

		err := templatex.Parse(`package "package"`).
			Save(file,
				templatex.WithFs(fsMock),
				templatex.WithFormat,
			)
		assert.ErrorIs(t, err, templatex.ErrFormatting)
	})
}
