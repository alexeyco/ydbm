package builder

import (
	"bytes"
	"sync"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"

	"github.com/alexeyco/ydbm/internal/columns"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// Builder describes query builder.
type Builder struct{}

// NewBuilder returns a new query buider.
func NewBuilder() *Builder {
	return &Builder{}
}

// CurrentVersion returns query to get current version.
func (b *Builder) CurrentVersion(path, name string) (string, *table.QueryParameters) {
	//nolint:forcetypeassert
	buf := bufferPool.Get().(*bytes.Buffer)

	buf.WriteString("select `")
	buf.WriteString(columns.Version)
	buf.WriteString("` from `")
	buf.WriteString(path)
	buf.WriteString("/")
	buf.WriteString(name)
	buf.WriteString("` order by `version` desc limit 1")

	q := buf.String()
	buf.Reset()
	bufferPool.Put(buf)

	return q, nil
}

// Insert returns insert query.
func (b *Builder) Insert(path, name string, migration Migration, appliedAt time.Time) (string, *table.QueryParameters) {
	//nolint:forcetypeassert
	buf := bufferPool.Get().(*bytes.Buffer)

	buf.WriteString("declare $v1 as int64; ")
	buf.WriteString("declare $v2 as string; ")
	buf.WriteString("declare $v3 as timestamp; ")

	buf.WriteString("insert into `")
	buf.WriteString(path)
	buf.WriteString("/")
	buf.WriteString(name)
	buf.WriteString("` (`")
	buf.WriteString(columns.Version)
	buf.WriteString("`,`")
	buf.WriteString(columns.Info)
	buf.WriteString("`,`")
	buf.WriteString(columns.AppliedAt)
	buf.WriteString("`) values ($v1,$v2,$v3)")

	q := buf.String()
	buf.Reset()
	bufferPool.Put(buf)

	return q, table.NewQueryParameters(
		table.ValueParam("$v1", types.Int64Value(migration.Version())),
		table.ValueParam("$v2", types.StringValueFromString(migration.Info())),
		table.ValueParam("$v3", types.TimestampValueFromTime(appliedAt)),
	)
}

// Delete returns delete query.
func (b *Builder) Delete(path, name string, migration Migration) (string, *table.QueryParameters) {
	//nolint:forcetypeassert
	buf := bufferPool.Get().(*bytes.Buffer)

	buf.WriteString("declare $v1 as int64; ")
	buf.WriteString("delete from `")
	buf.WriteString(path)
	buf.WriteString("/")
	buf.WriteString(name)
	buf.WriteString("` where `")
	buf.WriteString(columns.Version)
	buf.WriteString("`=$v1")

	q := buf.String()
	buf.Reset()
	bufferPool.Put(buf)

	return q, table.NewQueryParameters(
		table.ValueParam("$v1", types.Int64Value(migration.Version())),
	)
}
