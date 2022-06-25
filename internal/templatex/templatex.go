package templatex

import (
	"bytes"
	"fmt"
	"go/format"
	"text/template"

	"github.com/spf13/afero"
)

// Template describes template.
type Template struct {
	tpl *template.Template
}

// Parse returns new Template.
func Parse(s string) *Template {
	return &Template{
		tpl: template.Must(template.New("").Parse(s)),
	}
}

// Save compiles template and writes result to specified file.
func (t *Template) Save(file string, opts ...Option) error {
	o := Options{
		Fs: afero.NewOsFs(),
	}

	for _, opt := range opts {
		opt(&o)
	}

	var buf bytes.Buffer
	if err := t.tpl.Execute(&buf, o.Data); err != nil {
		return fmt.Errorf("%w: %v", ErrTemplate, err)
	}

	f, err := o.Fs.Create(file)
	if err != nil {
		return err
	}

	defer func() { _ = f.Close() }()

	b := buf.Bytes()
	if o.Format {
		if b, err = format.Source(b); err != nil {
			return fmt.Errorf("%w: %v", ErrFormatting, err)
		}
	}

	_, err = f.Write(b)

	return err
}
