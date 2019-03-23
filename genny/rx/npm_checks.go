package rx

import (
	"github.com/gobuffalo/genny"
)

func npmChecks(opts *Options) *genny.Generator {
	t := Tool{
		Name:    "NPM",
		Bin:     "npm",
		Minimum: []string{">=6.0.0", ">=7.0.0"},
		Partial: "node/_npm.plush",
		Version: func(r *genny.Runner) (string, error) {
			if v, ok := opts.Versions.Load("npm"); ok {
				return v, nil
			}
			return cmdVersion(r, "npm", "--version")
		},
	}

	g := t.Generator(opts)
	return g
}
