package timex

import "time"

// Clock describes mockable time struct.
type Clock struct{}

// New returns a new Clock instance.
func New() *Clock {
	return &Clock{}
}

// Now returns current time.
func (*Clock) Now() time.Time {
	return time.Now()
}
