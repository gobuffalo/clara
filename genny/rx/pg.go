package rx

import (
	"regexp"
	"strings"

	"github.com/gobuffalo/genny"
)

var PGMinimums = []string{">=9.3", ">=10.0", ">=11.0"}
var pgrx = regexp.MustCompile(`\d+\.\d+\.?\d*`)

func postgresChecks(opts *Options) *genny.Generator {
	t := Tool{
		Name:    "PostgreSQL",
		Bin:     "postgres",
		Minimum: PGMinimums,
		Partial: "db/_pg.plush",
		Version: func(r *genny.Runner) (string, error) {
			if v, ok := opts.Versions.Load("postgres"); ok {
				return v, nil
			}
			v, err := cmdVersion(r, "postgres", "--version")
			if err != nil {
				return "", err
			}
			v = pgrx.FindString(v)
			v = strings.TrimSpace(v)
			return v + ".0", nil
		},
	}

	g := t.Generator(opts)
	return g
}

// PostgreSQL (>= 9.3)
// CockroachDB (>= 1.1.1)
// MySQL (>= 5.7)
// SQLite3 (>= 3.x)
