package generator

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/afero"

	"github.com/alexeyco/ydbm/internal/generator/actions/executor"
	"github.com/alexeyco/ydbm/internal/generator/actions/migration"
	"github.com/alexeyco/ydbm/internal/generator/context"
)

var migrationFileRegex = regexp.MustCompile(`^\d+_[\w\d_]+\.go$`)

// Generator describes migrations generator.
type Generator struct {
	fs      afero.Fs
	actions []Action
	info    string
}

// Generate returns new Generator.
func Generate(info string, opts ...Option) *Generator {
	g := &Generator{
		fs: afero.NewOsFs(),
		actions: []Action{
			migration.New(),
			executor.New(),
		},
		info: info,
	}

	for _, opt := range opts {
		opt(g)
	}

	return g
}

// To generate a migration in specified directory.
func (g *Generator) To(directory string) error {
	if err := g.fs.MkdirAll(directory, os.ModePerm); err != nil {
		return err
	}

	current, err := g.current(directory)
	if err != nil {
		return err
	}

	var nextVersion int64 = 1
	if len(current) > 0 {
		nextVersion = current[len(current)-1].Version + 1
	}

	ctx := context.Context{
		Fs:        g.fs,
		Directory: directory,
		Current:   current,
		New: context.Migration{
			Version: nextVersion,
			Info:    g.info,
			Struct:  g.structName(nextVersion),
		},
	}

	for _, action := range g.actions {
		if err = action.Generate(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) current(directory string) ([]context.Migration, error) {
	var migrations []context.Migration

	err := afero.Walk(g.fs, directory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !migrationFileRegex.MatchString(info.Name()) {
			return nil
		}

		ver := strings.Split(info.Name(), "_")
		version, err := strconv.ParseInt(ver[0], 10, 64)
		if err != nil {
			return err
		}

		migrations = append(migrations, context.Migration{
			Version: version,
			Struct:  g.structName(version),
		})

		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return migrations, nil
}

func (g *Generator) structName(version int64) string {
	return fmt.Sprintf("migration%03d", version)
}
