package context

import "github.com/spf13/afero"

// Context describes Generator context.
type Context struct {
	Fs        afero.Fs
	Directory string
	Current   []Migration
	New       Migration
}

// Migration describes migration.
type Migration struct {
	Version int64
	Info    string
	Struct  string
}
