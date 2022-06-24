package executor

import "text/template"

// nolint:gci
var tpl = template.Must(template.New("").Parse(`package {{ .Package }}

import (
	"github.com/alexeyco/ydbm"
	"github.com/ydb-platform/ydb-go-sdk/v3"
)

// Executor returns migrations executor.
func Executor(conn ydb.Connection) *ydbm.Executor {
	return ydbm.New(conn).
		Register({{ range .Migrations }}
			{{ .Struct }}{},{{ end }}
		)
}
`))
