package rx

import (
	"regexp"
	"strings"

	"github.com/gobuffalo/genny/v2"
)

var MySQLMinimums = []string{">=3.5"}
var myrx = regexp.MustCompile(`\d+\.\d+\.?\d*`)

func mysqlChecks(opts *Options) *genny.Generator {
	t := Tool{
		Name:    "MySQL",
		Bin:     "mysql",
		Minimum: MySQLMinimums,
		Partial: "db/_mysql.plush",
		Version: func(r *genny.Runner) (string, error) {
			if v, ok := opts.Versions.Load("mysql"); ok {
				return v, nil
			}
			v, err := cmdVersion(r, "mysql", "--version")
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

// MySQL (>= 9.3)
// CockroachDB (>= 1.1.1)
// MySQL (>= 5.7)
// SQLite3 (>= 3.x)
