package ydbm

import (
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"

	"github.com/alexeyco/ydbm/internal/builder"
)

// Builder describes query builder interface.
type Builder interface {
	CurrentVersion(string, string) (string, *table.QueryParameters)
	Insert(string, string, builder.Migration, time.Time) (string, *table.QueryParameters)
	Delete(string, string, builder.Migration) (string, *table.QueryParameters)
}
