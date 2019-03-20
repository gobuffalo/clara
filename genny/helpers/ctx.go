package helpers

import (
	"io"
	"strings"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/plush"
)

func Context(opts interface{}) *plush.Context {
	ctx := plush.NewContext()
	ctx.Set("PATH", envy.Get("PATH", "PATH NOT FOUND"))
	ctx.Set("GOPATH", envy.GoPath())
	ctx.Set("opts", opts)
	ctx.Set("error", Error)
	ctx.Set("warning", Warning)
	ctx.Set("success", Success)
	ctx.Set("join", strings.Join)
	return ctx
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
