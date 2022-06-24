package templatex

import (
	"bytes"
	"go/format"
	"text/template"

	"github.com/spf13/afero"
)

// CompileWrite compiles template and writes result to a file.
func CompileWrite(fs afero.Fs, fileName string, tpl *template.Template, opts ...Option) error {
	o := options{
		header: true,
		format: true,
	}

	for _, opt := range opts {
		opt(&o)
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, o.data); err != nil {
		return err
	}

	f, err := fs.Create(fileName)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	b := buf.Bytes()

	if o.header {
		if _, err = f.WriteString(comment + "\n"); err != nil {
			return err
		}
	}

	if o.format {
		if b, err = format.Source(b); err != nil {
			return err
		}
	}

	_, err = f.Write(b)

	return err
}
