package templatex

type options struct {
	data   any
	header bool
	format bool
}

// Option describes option.
type Option func(*options)

// Data used to pass data to template.
func Data(data any) Option {
	return func(o *options) {
		o.data = data
	}
}

// WithoutHeader disables header in result file.
func WithoutHeader(o *options) {
	o.header = false
}

// DoNotFormat disables file formatting after template compilation.
func DoNotFormat(o *options) {
	o.format = false
}
