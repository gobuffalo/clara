package rx

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/genny/v2"
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
			if v, ok := opts.Versions.Load("go"); ok {
				return v, nil
			}
			v, err := cmdVersion(r, "go", "version")
			if err != nil {
				return "", err
			}
			v = strings.Split(v, " ")[2] // go version go1.17.8 linux/amd64
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
		ctx.Set("exec", "go")
		return opts.render("go/pkg_found.plush", ctx)
	}
}

func goPathBinCheck(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		opts.Out.Header("Go: Checking PATH")
		path := os.Getenv("PATH")

		ctx := Context(opts)
		if strings.Contains(path, filepath.Join("go", "bin")) {
			return opts.render("valid_path.plush", ctx)
		}
		return opts.render("invalid_path.plush", ctx)
	}
}
