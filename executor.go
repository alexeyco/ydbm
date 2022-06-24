package ydbm

import (
	"github.com/ydb-platform/ydb-go-sdk/v3"

	"github.com/alexeyco/ydbm/internal/generator/context"
)

// Executor describes migrations executor.
type Executor struct {
	conn       ydb.Connection
	migrations []Migration
}

// New returns new Executor.
func New(conn ydb.Connection) *Executor {
	return &Executor{
		conn: conn,
	}
}

// Register registers migrations to executor.
func (e *Executor) Register(migrations ...Migration) *Executor {
	e.migrations = append(e.migrations, migrations...)

	return e
}

// Up increments database version.
func (e *Executor) Up(ctx context.Context) error {
	return nil
}

// Down decrements database version.
func (e *Executor) Down(ctx context.Context) error {
	return nil
}
