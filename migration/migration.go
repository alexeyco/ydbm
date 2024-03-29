package migration

import (
	"context"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

// Migration describes migration contract.
type Migration interface {
	Version() int64
	Info() string
	Up(context.Context, string, table.Session) error
	Down(context.Context, string, table.Session) error
}
