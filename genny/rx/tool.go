package rx

import (
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/gobuffalo/genny"
)

type Tool struct {
	Name    string
	Bin     string
	Version func(*genny.Runner) (string, error)
	Partial string
	Minimum []string
}

func (t Tool) AcceptVersion(v string) (bool, error) {
	sv, err := semver.NewVersion(v)
	if err != nil {
		return false, err
	}
	for _, x := range t.Minimum {
		c, err := semver.NewConstraint(x)
		if err != nil {
			return false, err
		}
		if c.Check(sv) {
			return true, nil
		}
	}
	return false, nil
}

func (t Tool) Generator(opts *Options) *genny.Generator {
	g := genny.New()

	ctx := Context(opts)
	ctx.Set("tool", t)

	g.RunFn(func(r *genny.Runner) error {
		opts.Out.Header(fmt.Sprintf("%s: Checking installation", t.Name))
		bin, err := r.LookPath(t.Bin)
		if err != nil {
			x, err := templates.FindString("exec_not_found.plush")
			if err != nil {
				return opts.Out.RenderE(err)
			}
			return opts.Out.Render(x, ctx)
		}
		x, err := templates.FindString("exec_found.plush")
		if err != nil {
			return opts.Out.RenderE(err)
		}
		ctx.Set("bin", bin)
		return opts.Out.Render(x, ctx)
	})

	g.RunFn(func(r *genny.Runner) error {
		_, err := r.LookPath(t.Bin)
		if err != nil {
			return nil
		}
		opts.Out.Header(fmt.Sprintf("%s: Checking minimum version requirements", t.Name))
		v, err := t.Version(r)
		if err != nil {
			return opts.Out.RenderE(err)
		}
		ctx.Set("version", v)
		b, err := t.AcceptVersion(v)
		if err != nil {
			return opts.Out.RenderE(err)
		}
		if b {
			return opts.render("min_version.plush", ctx)
		}
		return opts.render("non_min_version.plush", ctx)
	})
	return g
}
