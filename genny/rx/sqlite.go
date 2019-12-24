package rx

import (
	"regexp"
	"strings"

	"github.com/gobuffalo/genny/v2"
)

var SQLiteMinimums = []string{">=3.0"}
var sqliterx = regexp.MustCompile(`^\d+\.\d+\.?\d*`)

func sqliteChecks(opts *Options) *genny.Generator {
	t := Tool{
		Name:    "SQLite3",
		Bin:     "sqlite3",
		Minimum: SQLiteMinimums,
		Partial: "db/_sqlite.plush",
		Version: func(r *genny.Runner) (string, error) {
			if v, ok := opts.Versions.Load("sqlite3"); ok {
				return v, nil
			}
			v, err := cmdVersion(r, "sqlite3", "--version")
			if err != nil {
				return "", err
			}
			v = sqliterx.FindString(v)
			v = strings.TrimSpace(v)
			return v, nil
		},
	}

	g := t.Generator(opts)
	return g
}

// SQLite (>= 9.3)
// CockroachDB (>= 1.1.1)
// SQLite (>= 5.7)
// SQLite3 (>= 3.x)
