package ydbm

import "time"

// Clock describes clock interface.
type Clock interface {
	Now() time.Time
}
