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

// WithActions to set custom actions.
func WithActions(a ...Action) Option {
	return func(g *Generator) {
		g.actions = a
	}
}
