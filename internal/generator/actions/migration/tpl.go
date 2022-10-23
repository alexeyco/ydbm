package migration

import "github.com/alexeyco/ydbm/internal/templatex"

//nolint:gci
var tpl = templatex.Parse(`package {{ .Package }}

import (
	"context"

	"github.com/alexeyco/ydbm/migration"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

var _ migration.Migration = (*{{ .Struct }})(nil)

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
func ({{ .Struct }}) Up(ctx context.Context, path string, s table.Session) error {
	return nil
}

// Down to decrement database version.
func ({{ .Struct }}) Down(ctx context.Context, path string, s table.Session) error {
	return nil
}
`)
