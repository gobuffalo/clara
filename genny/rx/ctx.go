package rx

import (
	"os"
	"strings"

	"github.com/gobuffalo/plush"
)

func Context(opts interface{}) *plush.Context {
	var h Helpers
	ctx := plush.NewContext()
	ctx.Set("PATH", os.Getenv("PATH"))
	ctx.Set("GOPATH", os.Getenv("GOPATH"))
	ctx.Set("opts", opts)
	ctx.Set("error", h.Error)
	ctx.Set("warning", h.Warning)
	ctx.Set("success", h.Success)
	ctx.Set("join", strings.Join)
	return ctx
}
