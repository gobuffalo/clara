package rx

import (
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/here/there"
)

var GoMinimums = []string{">=1.13"}

func goCheck(opts *Options) *genny.Generator {
	t := Tool{
		Name:    "Go",
		Bin:     "go",
		Minimum: GoMinimums,
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
	g.RunFn(goPkgCheck(opts))
	g.RunFn(goPathBinCheck(opts))
	return g
}

func goPkgCheck(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		opts.Out.Header("Go: Checking Package Management")
		ctx := Context(opts)

		info, err := there.Current()
		if err != nil {
			return err
		}
		if len(info.Module.Path) == 0 {
			return opts.render("go/pkg_not_found.plush", ctx)
		}
		ctx.Set("pkg", "Go Modules")
		ctx.Set("exec", envy.GoBin())
		return opts.render("go/pkg_found.plush", ctx)
	}
}

func goPathBinCheck(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		opts.Out.Header("Go: Checking PATH")
		path := envy.Get("PATH", "")

		ctx := Context(opts)
		if strings.Contains(path, filepath.Join(envy.GoPath(), "bin")) {
			return opts.render("valid_path.plush", ctx)
		}
		return opts.render("invalid_path.plush", ctx)
	}
}
