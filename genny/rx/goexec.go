package rx

import (
	"github.com/gobuffalo/doctor/genny/helpers"
	"github.com/gobuffalo/genny"
	"github.com/pkg/errors"
)

func goExec(opts *Options) genny.RunFn {
	ctx := helpers.Context(opts)
	return func(r *genny.Runner) error {
		helpers.Header(opts.Out, "Checking for the Go executable")
		bin, err := r.LookPath("go")
		if err != nil {
			x, err := templates.FindString("go_not_found.plush")
			if err != nil {
				return errors.WithStack(err)
			}
			return helpers.Render(opts.Out, x, ctx)
		}
		x, err := templates.FindString("go_found.plush")
		if err != nil {
			return errors.WithStack(err)
		}
		ctx.Set("bin", bin)
		return helpers.Render(opts.Out, x, ctx)
	}
}
