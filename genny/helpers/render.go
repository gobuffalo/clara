package helpers

import (
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
	"github.com/gobuffalo/plush"
)

func RenderE(w io.Writer, err error) error {
	s := color.RedString(fmt.Sprintf("%s %v", ERROR, err))
	s = strings.TrimSpace(s)
	s += "\n\n"
	w.Write([]byte(s))
	return nil
}

func Render(w io.Writer, s string, ctx *plush.Context) error {
	s, err := plush.Render(s, ctx)
	if err != nil {
		return err
	}
	s = strings.TrimSpace(s)
	s += "\n\n"
	_, err = w.Write([]byte(s))
	return err
}
