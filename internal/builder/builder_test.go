package builder_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"

	"github.com/alexeyco/ydbm/internal/builder"
)

const (
	path    = "/path/to"
	name    = "migrations"
	version = int64(123)
	info    = "info"
)

var appliedAt = time.Now()

func TestBuilder_CurrentVersion(t *testing.T) {
	t.Parallel()

	b := builder.NewBuilder()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		expectedQuery := "select `version` from `/path/to/migrations` order by `version` desc limit 1"

		actualQuery, actualParams := b.CurrentVersion(path, name)

		assert.Equal(t, expectedQuery, actualQuery)
		assert.Nil(t, actualParams)
	})
}

func TestBuilder_Insert(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	migrationMock := NewMockMigration(ctrl)

	b := builder.NewBuilder()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		expectedQuery := "declare $v1 as int64; " +
			"declare $v2 as string; " +
			"declare $v3 as timestamp; " +
			"insert into `/path/to/migrations` (`version`,`info`,`applied_at`) values ($v1,$v2,$v3)"

		expectedParams := table.NewQueryParameters(
			table.ValueParam("$v1", types.Int64Value(version)),
			table.ValueParam("$v2", types.StringValueFromString(info)),
			table.ValueParam("$v3", types.TimestampValueFromTime(appliedAt)),
		)

		migrationMock.EXPECT().Version().Return(version)
		migrationMock.EXPECT().Info().Return(info)

		actualQuery, actualParams := b.Insert(path, name, migrationMock, appliedAt)

		assert.Equal(t, expectedQuery, actualQuery)
		assert.Equal(t, expectedParams, actualParams)
	})
}

func TestBuilder_Delete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	migrationMock := NewMockMigration(ctrl)

	b := builder.NewBuilder()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		expectedQuery := "declare $v1 as int64; " +
			"delete from `/path/to/migrations` where `version`=$v1"

		expectedParams := table.NewQueryParameters(
			table.ValueParam("$v1", types.Int64Value(version)),
		)

		migrationMock.EXPECT().Version().Return(version)

		actualQuery, actualParams := b.Delete(path, name, migrationMock)

		assert.Equal(t, expectedQuery, actualQuery)
		assert.Equal(t, expectedParams, actualParams)
	})
}
