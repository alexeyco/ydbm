package templatex

import "github.com/spf13/afero"

// Options describes options.
type Options struct {
	Fs     afero.Fs
	Format bool
	Data   map[string]any
}

// Option describes template option.
type Option func(*Options)

// WithFs sets custom fs.
func WithFs(fs afero.Fs) Option {
	return func(o *Options) {
		o.Fs = fs
	}
}

// WithFormat enables formatting after writing.
func WithFormat(o *Options) {
	o.Format = true
}

// WithData sets template data.
func WithData(data map[string]any) Option {
	return func(o *Options) {
		o.Data = data
	}
}
