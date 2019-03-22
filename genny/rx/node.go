package rx

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/gobuffalo/clara/genny/helpers"
	"github.com/gobuffalo/genny"
)

func checkNode(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		helpers.Header(opts.Out, "Checking Node")
		ctx := helpers.Context(opts)
		if _, err := r.LookPath("node"); err != nil {
			return opts.render("missing_node.plush", ctx)
		}
		bb := &bytes.Buffer{}
		c := exec.Command("node", "--version")
		c.Stdout = bb
		c.Stderr = bb
		if err := r.Exec(c); err != nil {
			return err
		}
		v, err := semver.NewVersion(strings.TrimSpace(bb.String()))
		if err != nil {
			return err
		}
		ctx.Set("node", v)
		if v.Major() < 11 {
			return opts.render("invalid_node_version.plush", ctx)
		}
		return opts.render("valid_node.plush", ctx)
	}
}

func checkYarn(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		helpers.Header(opts.Out, "Checking YARN")
		ctx := helpers.Context(opts)
		if _, err := r.LookPath("yarn"); err != nil {
			return opts.render("missing_yarn.plush", ctx)
		}
		bb := &bytes.Buffer{}
		c := exec.Command("yarn", "--version")
		c.Stdout = bb
		c.Stderr = bb
		if err := r.Exec(c); err != nil {
			return helpers.Render(opts.Out, bb.String(), ctx)
		}
		v, err := semver.NewVersion(strings.TrimSpace(bb.String()))
		if err != nil {
			return err
		}
		ctx.Set("yarn", v)
		if v.Major() < 1 {
			return opts.render("invalid_yarn_version.plush", ctx)
		}
		if v.Minor() < 12 {
			return opts.render("invalid_yarn_version.plush", ctx)
		}
		return opts.render("valid_yarn.plush", ctx)
	}
}

func checkNpm(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		helpers.Header(opts.Out, "Checking NPM")
		ctx := helpers.Context(opts)
		if _, err := r.LookPath("npm"); err != nil {
			return opts.render("missing_npm.plush", ctx)
		}
		bb := &bytes.Buffer{}
		c := exec.Command("npm", "--version")
		c.Stdout = bb
		c.Stderr = bb
		if err := r.Exec(c); err != nil {
			return helpers.Render(opts.Out, bb.String(), ctx)
		}
		v, err := semver.NewVersion(strings.TrimSpace(bb.String()))
		if err != nil {
			return err
		}
		ctx.Set("npm", v)
		if v.Major() < 6 {
			return opts.render("invalid_npm_version.plush", ctx)
		}
		if v.Minor() < 7 {
			return opts.render("invalid_npm_version.plush", ctx)
		}
		return opts.render("valid_npm.plush", ctx)
	}
}
