package helpers

import (
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
