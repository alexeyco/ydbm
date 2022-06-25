package generator

import "github.com/spf13/afero"

// Option describes Generator option.
type Option func(*Generator)

// WithFs to set custom Fs.
func WithFs(fs afero.Fs) Option {
	return func(g *Generator) {
		g.fs = fs
	}
}

// WithValidator to set custom Validator.
func WithValidator(v Validator) Option {
	return func(g *Generator) {
		g.validator = v
	}
}

// WithActions to set custom actions.
func WithActions(a ...Action) Option {
	return func(g *Generator) {
		g.actions = a
	}
}
