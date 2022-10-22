package builder

// Migration describes migration interface.
type Migration interface {
	Version() int64
	Info() string
}
