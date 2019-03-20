package rx

import (
	"strings"

	"github.com/Masterminds/semver"
	"github.com/gobuffalo/doctor/genny/helpers"
	"github.com/gobuffalo/genny"
	"github.com/pkg/errors"
)

func goVersion(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		helpers.Header(opts.Out, "Checking your Go version")
		v := opts.GoVersion
		v = strings.TrimPrefix(v, "go")
		sv, err := semver.NewVersion(v)
		if err != nil {
			return errors.WithStack(err)
		}

		ctx := helpers.Context(opts)
		ctx.Set("version", v)
		min := func() error {
			return opts.render("non_min_go.plush", ctx)
		}
		switch sv.Minor() {
		case 12:
			return opts.render("min_go.plush", ctx)
		case 10:
			if sv.Patch() < 8 {
				return min()
			}
		case 11:
			if sv.Patch() < 4 {
				return min()
			}
		default:
			return min()
		}
		return nil
	}
}
