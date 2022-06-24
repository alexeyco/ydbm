package logx

import (
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/hashicorp/go-multierror"
)

var (
	info = color.Style{
		color.Bold,
		color.HiGreen,
	}

	fatal = color.Style{
		color.Red,
	}

	fatalBold = color.Style{
		color.Bold,
		color.HiRed,
	}
)

// Logger describes logger.
type Logger struct{}

// New returns new Logger.
func New() *Logger {
	return &Logger{}
}

// Fatal prints fatal error.
func (l *Logger) Fatal(err error) {
	// nolint:errorlint
	merr, ok := err.(*multierror.Error)
	if ok {
		l.line(fatalBold, "Following errors occurred:")

		for _, e := range merr.Errors {
			l.line(fatal, "  * %v", e)
		}
	} else {
		l.line(fatalBold, err.Error())
	}

	os.Exit(0)
}

// Infof prints info message.
func (l *Logger) Infof(format string, a ...any) {
	l.line(info, format, a...)
}

func (l *Logger) line(st color.Style, format string, a ...any) {
	st.Println(fmt.Sprintf(format, a...))
}
