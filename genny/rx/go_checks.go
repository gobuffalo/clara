package rx

import (
	"go/build"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gobuffalo/clara/genny/helpers"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/genny"
)

func goCheck(opts *Options) *genny.Generator {
	t := Tool{
		Name:    "Go",
		Bin:     "go",
		Minimum: []string{">=1.10.8", ">=1.11.4", ">=1.12"},
		Partial: "go/_help.plush",
		Version: func(r *genny.Runner) (string, error) {
			v, ok := opts.Versions.Load("go")
			if !ok {
				v = runtime.Version()
			}
			return strings.TrimPrefix(v, "go"), nil
		},
	}
	g := t.Generator(opts)
	g.RunFn(goPathCheck(opts))
	g.RunFn(goPkgCheck(opts))
	g.RunFn(goPathBinCheck(opts))
	return g
}

func goPathCheck(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		ctx := helpers.Context(opts)
		helpers.Header(opts.Out, "Go: Checking GOPATH")
		if envy.Mods() {
			return opts.render("go/using_mods.plush", ctx)
		}
		src := filepath.Join(envy.GoPath(), "src")
		if strings.HasPrefix(opts.App.Pwd, src) {
			return opts.render("go/good_gopath.plush", ctx)
		}
		c := build.Default
		ctx.Set("sources", c.SrcDirs())
		return opts.render("go/bad_gopath.plush", ctx)
	}
}

func goPkgCheck(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		helpers.Header(opts.Out, "Go: Checking Package Management")
		ctx := helpers.Context(opts)
		if envy.Mods() {
			ctx.Set("pkg", "Go Modules")
			ctx.Set("exec", envy.GoBin())
			return opts.render("go/pkg_found.plush", ctx)
		}
		if opts.App.WithDep {
			ex, err := r.LookPath("dep")
			if err != nil {
				return opts.render("go/dep_not_found.plush", ctx)
			}
			ctx.Set("pkg", "Dep")
			ctx.Set("exec", ex)
			return opts.render("go/pkg_found.plush", ctx)
		}
		return opts.render("go/pkg_not_found.plush", ctx)
	}
}

func goPathBinCheck(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		helpers.Header(opts.Out, "Go: Checking PATH")
		path := envy.Get("PATH", "")

		ctx := helpers.Context(opts)
		if strings.Contains(path, filepath.Join(envy.GoPath(), "bin")) {
			return opts.render("valid_path.plush", ctx)
		}
		return opts.render("invalid_path.plush", ctx)
	}
}
