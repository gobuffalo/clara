package rx

import (
	"strings"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/plush"
)

func Context(opts interface{}) *plush.Context {
	var h Helpers
	ctx := plush.NewContext()
	ctx.Set("PATH", envy.Get("PATH", "PATH NOT FOUND"))
	ctx.Set("GOPATH", envy.GoPath())
	ctx.Set("opts", opts)
	ctx.Set("error", h.Error)
	ctx.Set("warning", h.Warning)
	ctx.Set("success", h.Success)
	ctx.Set("join", strings.Join)
	return ctx
}
