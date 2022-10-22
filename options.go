package ydbm

// Option describes executor option.
type Option func(*Executor)

// WithTablePath sets custom table path.
func WithTablePath(path string) Option {
	return func(e *Executor) {
		e.path = path
	}
}

// WithTableName sets custom table name.
func WithTableName(name string) Option {
	return func(e *Executor) {
		e.name = name
	}
}
