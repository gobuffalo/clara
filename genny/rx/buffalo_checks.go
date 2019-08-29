package rx

import (
	"regexp"
	"strings"

	"github.com/gobuffalo/genny"
)

var bvrx = regexp.MustCompile(`v\d+\.\d+\.\d+`)
var BuffaloMinimums = []string{">=0.14.10"}

func buffaloChecks(opts *Options) *genny.Generator {
	t := Tool{
		Name:    "Buffalo (CLI)",
		Bin:     "buffalo",
		Minimum: BuffaloMinimums,
		Partial: "buffalo/_help.plush",
		Version: func(r *genny.Runner) (string, error) {
			if v, ok := opts.Versions.Load("buffalo"); ok {
				return v, nil
			}

			v, err := cmdVersion(r, "buffalo", "version")
			if err != nil {
				return v, err
			}
			v = bvrx.FindString(v)
			v = strings.TrimSpace(v)
			return v, nil

		},
	}
	g := t.Generator(opts)
	return g
}
