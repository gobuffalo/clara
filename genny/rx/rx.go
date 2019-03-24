package rx

import (
	"github.com/gobuffalo/genny"
	packr "github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

var templates = packr.New("github.com/gobuffalo/clara/genny/rx/templates", "../rx/templates")

func init() {
	plush.Helpers.Add("partialFeeder", templates.FindString)
}

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, errors.WithStack(err)
	}

	g.Merge(goCheck(opts))
	if !opts.SkipBuffalo {
		g.Merge(buffaloChecks(opts))
		if !opts.SkipNode {
			g.Merge(nodeChecks(opts))
			g.Merge(npmChecks(opts))
			g.Merge(yarnChecks(opts))
		}
	}
	return g, nil
}
