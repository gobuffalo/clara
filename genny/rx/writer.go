package rx

import (
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
)

type Writer struct {
	io.Writer
}

func NewWriter(w io.Writer) Writer {
	ww := Writer{
		Writer: w,
	}
	return ww
}

func (w Writer) WriteString(s string) error {
	_, err := w.Write([]byte(s))
	return err
}

func (w Writer) Header(s string) {
	s = strings.TrimSpace(s)
	s = fmt.Sprintf("-> %s\n", s)
	w.Write([]byte(s))
}

func (w Writer) Success(s string) {
	w.colorize(SUCCESS, s)
}

func (w Writer) Error(s string) {
	w.colorize(ERROR, s)
}

func (w Writer) Warning(s string) {
	w.colorize(WARNING, s)
}

func (w Writer) colorize(p string, x string) {
	x = strings.TrimSpace(x)
	x = fmt.Sprintf("%s %s", p, x)
	switch p {
	case SUCCESS:
		w.WriteString(color.GreenString(x))
	case ERROR:
		w.WriteString(color.RedString(x))
	case WARNING:
		w.WriteString(color.YellowString(x))
	}
}
