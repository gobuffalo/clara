package rx

import (
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
	"github.com/gobuffalo/plush/v4"
)

func (w Writer) RenderE(err error) error {
	s := color.RedString(fmt.Sprintf("%s %v", ERROR, err))
	s = strings.TrimSpace(s)
	s += "\n\n"
	return w.WriteString(s)
}

func (w Writer) Render(s string, ctx *plush.Context) error {
	ctx.Set("partialFeeder", templateFeeder)
	s, err := plush.Render(s, ctx)
	if err != nil {
		return err
	}
	s = strings.TrimSpace(s)
	s += "\n\n"
	return w.WriteString(s)
}

func templateFeeder(name string) (s string, err error) {
	s = ""
	f, err := templates.Open("templates/" + name)
	if err != nil {
		return
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return
	}
	s = string(b)
	return
}
