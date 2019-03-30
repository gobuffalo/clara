package rx

import (
	"regexp"
	"strings"

	"github.com/gobuffalo/genny"
)

var CockroachMinimums = []string{">=1.1.1", ">=2.0.x"}
var crrx = regexp.MustCompile(`\d+\.\d+\.?\d*`)

func cockroachChecks(opts *Options) *genny.Generator {
	t := Tool{
		Name:    "Cockroach",
		Bin:     "cockroach",
		Minimum: CockroachMinimums,
		Partial: "db/_cockroach.plush",
		Version: func(r *genny.Runner) (string, error) {
			if v, ok := opts.Versions.Load("cockroach"); ok {
				return v, nil
			}
			v, err := cmdVersion(r, "cockroach", "version")
			if err != nil {
				return "", err
			}
			v = pgrx.FindString(v)
			v = strings.TrimSpace(v)
			return v, nil
		},
	}

	g := t.Generator(opts)
	return g
}

// cockroachQL (>= 9.3)
// CockroachDB (>= 1.1.1)
// MySQL (>= 5.7)
// SQLite3 (>= 3.x)
