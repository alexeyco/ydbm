package executor

import (
	"path"

	"github.com/alexeyco/ydbm/internal/generator/context"
	"github.com/alexeyco/ydbm/internal/templatex"
)

// Action describes action to generate migrations executor constructor.
type Action struct{}

// New returns new Action.
func New() *Action {
	return &Action{}
}

// Generate generates migrations executor constructor.
func (a *Action) Generate(ctx context.Context) error {
	packageName := path.Base(ctx.Directory)

	migrations := ctx.Current
	migrations = append(migrations, ctx.New)

	return tpl.Save(path.Join(ctx.Directory, "executor.go"),
		templatex.WithFs(ctx.Fs),
		templatex.WithFormat,
		templatex.WithData(map[string]any{
			"Package":    packageName,
			"Migrations": migrations,
		}))
}
