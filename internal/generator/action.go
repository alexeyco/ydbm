package generator

import (
	"github.com/alexeyco/ydbm/internal/generator/context"
)

// Action describes generator action.
type Action interface {
	Generate(context.Context) error
}
