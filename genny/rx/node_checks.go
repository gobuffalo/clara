package rx

import (
	"github.com/gobuffalo/genny"
)

func nodeChecks(opts *Options) *genny.Generator {
	t := Tool{
		Name:    "Node",
		Bin:     "node",
		Minimum: []string{">=1.11"},
		Partial: "node/_help.plush",
		Version: func(r *genny.Runner) (string, error) {
			if v, ok := opts.Versions.Load("node"); ok {
				return v, nil
			}
			return cmdVersion(r, "node", "--version")
		},
	}

	g := t.Generator(opts)
	return g
}
