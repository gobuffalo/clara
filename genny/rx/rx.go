package rx

import (
	"github.com/gobuffalo/genny"
	packr "github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
)

var templates = packr.New("github.com/gobuffalo/clara/genny/rx/templates", "../rx/templates")

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, errors.WithStack(err)
	}

	g.RunFn(goExec(opts))
	g.RunFn(goVersion(opts))
	g.RunFn(pkgManagement(opts))
	g.RunFn(gopath(opts))
	g.RunFn(checkPath(opts))
	g.RunFn(checkNode(opts))
	g.RunFn(checkNpm(opts))
	g.RunFn(checkYarn(opts))
	return g, nil
}
