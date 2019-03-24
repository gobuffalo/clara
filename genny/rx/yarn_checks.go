package rx

import (
	"github.com/gobuffalo/genny"
)

var YarnMinimums = []string{">=1.12"}

func yarnChecks(opts *Options) *genny.Generator {
	t := Tool{
		Name:    "Yarn",
		Bin:     "yarnpkg",
		Minimum: YarnMinimums,
		Partial: "node/_yarn.plush",
		Version: func(r *genny.Runner) (string, error) {
			if v, ok := opts.Versions.Load("yarn"); ok {
				return v, nil
			}
			return cmdVersion(r, "yarnpkg", "--version")
		},
	}

	g := t.Generator(opts)
	return g
}
