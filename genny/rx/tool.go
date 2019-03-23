package rx

import (
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/gobuffalo/clara/genny/helpers"
	"github.com/gobuffalo/genny"
	"github.com/pkg/errors"
)

type Tool struct {
	Name    string
	Bin     string
	Version func() (string, error)
	Partial string
	Minimum []string
}

func (t Tool) AcceptVersion(v string) (bool, error) {
	sv, err := semver.NewVersion(v)
	if err != nil {
		return false, errors.WithMessage(err, v)
	}
	for _, x := range t.Minimum {
		c, err := semver.NewConstraint(x)
		if err != nil {
			return false, errors.WithMessage(err, x)
		}
		if c.Check(sv) {
			return true, nil
		}
	}
	return false, nil
}

func (t Tool) Generator(opts *Options) *genny.Generator {
	g := genny.New()

	ctx := helpers.Context(opts)
	ctx.Set("tool", t)

	g.RunFn(func(r *genny.Runner) error {
		helpers.Header(opts.Out, fmt.Sprintf("%s: Checking installation", t.Name))
		bin, err := r.LookPath(t.Bin)
		if err != nil {
			x, err := templates.FindString("exec_not_found.plush")
			if err != nil {
				return helpers.RenderE(opts.Out, err)
			}
			return helpers.Render(opts.Out, x, ctx)
		}
		x, err := templates.FindString("exec_found.plush")
		if err != nil {
			return helpers.RenderE(opts.Out, err)
		}
		ctx.Set("bin", bin)
		return helpers.Render(opts.Out, x, ctx)
	})

	g.RunFn(func(r *genny.Runner) error {
		helpers.Header(opts.Out, fmt.Sprintf("%s: Checking minimum version requirements", t.Name))
		v, err := t.Version()
		if err != nil {
			return helpers.RenderE(opts.Out, errors.WithMessage(err, v))
		}
		ctx.Set("version", v)
		b, err := t.AcceptVersion(v)
		if err != nil {
			return helpers.RenderE(opts.Out, errors.WithMessage(err, v))
		}
		if b {
			return opts.render("min_version.plush", ctx)
		}
		return opts.render("non_min_version.plush", ctx)
	})
	return g
}
