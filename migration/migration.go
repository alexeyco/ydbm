package migration

import "github.com/ydb-platform/ydb-go-sdk/v3/table"

// Migration describes migration contract.
type Migration interface {
	Version() int64
	Info() string
	Up(table.TransactionActor) error
	Down(table.TransactionActor) error
}
