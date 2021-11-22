package rx

import (
	"embed"

	"github.com/gobuffalo/genny/v2"
)

//go:embed templates templates/*/*.plush
var templates embed.FS

// func init() {
// 	plush.Helpers.Add("partialFeeder", templates.FindString)
// }

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, err
	}

	g.Merge(goCheck(opts))
	if opts.SkipBuffalo {
		return g, nil
	}
	if !opts.SkipNode {
		g.Merge(nodeChecks(opts))
		g.Merge(npmChecks(opts))
		g.Merge(yarnChecks(opts))
	}

	if !opts.SkipDB {
		g.Merge(postgresChecks(opts))
		g.Merge(mysqlChecks(opts))
		g.Merge(sqliteChecks(opts))
		g.Merge(cockroachChecks(opts))
	}

	g.Merge(buffaloChecks(opts))

	return g, nil
}
