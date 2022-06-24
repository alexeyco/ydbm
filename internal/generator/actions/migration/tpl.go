package migration

import "text/template"

// nolint:gci
var tpl = template.Must(template.New("").Parse(`package {{ .Package }}

import (
	"github.com/alexeyco/ydbm"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

var _ ydbm.Migration = (*{{ .Struct }})(nil)

// {{ .Struct }} describes migration.
type {{ .Struct }} struct {}

// Version returns migration version.
func ({{ .Struct }}) Version() int64 {
	return {{ .Version }}
}

// Info returns migration info.
func ({{ .Struct }}) Info() string {
	return {{ .Info | printf "%q" }}
}

// Up to increment database version.
func ({{ .Struct }}) Up(tx table.TransactionActor) error {
	return nil
}

// Down to decrement database version.
func ({{ .Struct }}) Down(tx table.TransactionActor) error {
	return nil
}
`))
