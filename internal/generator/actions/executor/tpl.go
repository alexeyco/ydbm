package executor

import "github.com/alexeyco/ydbm/internal/templatex"

//nolint:gci
var tpl = templatex.Parse(`// Code generated by github.com/alexeyco/ydbm. DO NOT EDIT.
package {{ .Package }}

import (
	"github.com/alexeyco/ydbm"
	"github.com/ydb-platform/ydb-go-sdk/v3"
)

// Executor returns migrations executor.
func Executor(conn ydb.Connection, path, name string) *ydbm.Executor {
	return ydbm.New(conn, ydbm.WithTablePath(path), ydbm.WithTableName(name)).
		Register({{ range .Migrations }}
			{{ .Struct }}{},{{ end }}
		)
}
`)
