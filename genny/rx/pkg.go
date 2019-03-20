package rx

import (
	"github.com/gobuffalo/doctor/genny/helpers"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/genny"
)

func pkgManagement(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		helpers.Header(opts.Out, "Checking Package Management")
		ctx := helpers.Context(opts)
		if envy.Mods() {
			ctx.Set("pkg", "Go Modules")
			ctx.Set("exec", envy.GoBin())
			return opts.render("pkg_found.plush", ctx)
		}
		if opts.App.WithDep {
			ex, err := r.LookPath("dep")
			if err != nil {
				return opts.render("dep_not_found.plush", ctx)
			}
			ctx.Set("pkg", "Dep")
			ctx.Set("exec", ex)
			return opts.render("pkg_found.plush", ctx)
		}
		return opts.render("pkg_not_found.plush", ctx)
	}
}
