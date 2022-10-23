package migration

import (
	"fmt"
	"path"

	"github.com/gobeam/stringy"

	"github.com/alexeyco/ydbm/internal/generator/context"
	"github.com/alexeyco/ydbm/internal/templatex"
)

// Action provides action.
type Action struct{}

// New returns new Action.
func New() *Action {
	return &Action{}
}

// Generate migration.
func (a *Action) Generate(ctx context.Context) error {
	fileName := fmt.Sprintf("%04d_%s.go", ctx.New.Version, stringy.New(ctx.New.Info).SnakeCase().ToLower())
	packageName := path.Base(ctx.Directory)

	return tpl.Save(path.Join(ctx.Directory, fileName),
		templatex.WithFs(ctx.Fs),
		templatex.WithFormat,
		templatex.WithData(map[string]any{
			"Package": packageName,
			"Struct":  ctx.New.Struct,
			"Version": ctx.New.Version,
			"Info":    ctx.New.Info,
		}))
}
