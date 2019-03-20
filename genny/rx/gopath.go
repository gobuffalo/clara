package rx

import (
	"go/build"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/doctor/genny/helpers"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/genny"
)

func gopath(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		ctx := helpers.Context(opts)
		helpers.Header(opts.Out, "Checking GOPATH")
		if envy.Mods() {
			return opts.render("using_mods.plush", ctx)
		}
		src := filepath.Join(envy.GoPath(), "src")
		if strings.HasPrefix(opts.App.Pwd, src) {
			return opts.render("valid_gopath.plush", ctx)
		}
		c := build.Default
		ctx.Set("sources", c.SrcDirs())
		return opts.render("invalid_gopath.plush", ctx)
	}
}
