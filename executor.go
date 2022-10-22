package ydbm

import (
	"context"
	"sort"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/options"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result/named"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"

	"github.com/alexeyco/ydbm/internal/builder"
	"github.com/alexeyco/ydbm/internal/columns"
	"github.com/alexeyco/ydbm/internal/timex"
	"github.com/alexeyco/ydbm/migration"
)

var read = table.TxControl(
	table.BeginTx(
		table.WithOnlineReadOnly(),
	),
	table.CommitTx(),
)

// Executor describes migrations executor.
type Executor struct {
	conn       ydb.Connection
	path       string
	name       string
	builder    Builder
	clock      Clock
	migrations []migration.Migration
}

// New returns new Executor.
func New(conn ydb.Connection, opts ...Option) *Executor {
	e := &Executor{
		conn:    conn,
		path:    "/",
		builder: builder.NewBuilder(),
		clock:   timex.New(),
		name:    "migrations",
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

// Register registers migrations to executor.
func (e *Executor) Register(migrations ...migration.Migration) *Executor {
	e.migrations = append(e.migrations, migrations...)

	sort.Slice(e.migrations, func(i, j int) bool {
		return e.migrations[i].Version() < e.migrations[j].Version()
	})

	return e
}

// Up increments database version.
func (e *Executor) Up(ctx context.Context) (err error) {
	current, err := e.currentVersion(ctx)
	if err != nil {
		return
	}

	for _, m := range e.migrations {
		m := m

		if m.Version() <= current {
			continue
		}

		err = e.conn.Table().DoTx(ctx, func(ctx context.Context, tx table.TransactionActor) (err error) {
			if err = m.Up(tx); err != nil {
				return
			}

			query, params := e.builder.Insert(e.path, e.name, m, e.clock.Now())
			_, err = tx.Execute(ctx, query, params)

			return
		})
		if err != nil {
			return err
		}
	}

	return
}

// Down decrements database version.
func (e *Executor) Down(ctx context.Context) (err error) {
	current, err := e.currentVersion(ctx)
	if err != nil {
		return
	}

	for i := len(e.migrations) - 1; i >= 0; i-- {
		m := e.migrations[i]

		if m.Version() > current {
			continue
		}

		err = e.conn.Table().DoTx(ctx, func(ctx context.Context, tx table.TransactionActor) (err error) {
			if err = m.Down(tx); err != nil {
				return
			}

			query, params := e.builder.Delete(e.path, e.name, m)
			_, err = tx.Execute(ctx, query, params)

			return
		})
		if err != nil {
			return err
		}
	}

	return
}

func (e *Executor) currentVersion(ctx context.Context) (version int64, err error) {
	if err = e.prepareTable(ctx); err != nil {
		return
	}

	err = e.conn.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
		query, params := e.builder.CurrentVersion(e.path, e.name)

		_, res, err := s.Execute(ctx, read, query, params)
		if err != nil {
			return err
		}
		defer func() { _ = res.Close() }()

		if res.ResultSetCount() == 0 {
			return nil
		}

		for res.NextResultSet(ctx) {
			for res.NextRow() {
				if err = res.ScanNamed(named.Optional(columns.Version, &version)); err != nil {
					return err
				}
			}
		}

		return res.Err()
	})

	return
}

func (e *Executor) prepareTable(ctx context.Context) (err error) {
	err = e.conn.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
		_, err := s.DescribeTable(ctx, e.path)
		if err == nil {
			return nil
		}

		if !ydb.IsOperationErrorSchemeError(err) {
			return err
		}

		err = s.CreateTable(ctx, e.path,
			options.WithColumn(columns.Version, types.Optional(types.TypeInt64)),
			options.WithColumn(columns.Info, types.Optional(types.TypeString)),
			options.WithColumn(columns.AppliedAt, types.Optional(types.TypeTimestamp)),
			options.WithPrimaryKeyColumn(columns.Version),
		)

		return err
	})

	return
}
