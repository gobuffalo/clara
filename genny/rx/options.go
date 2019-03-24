package rx

import (
	"io"
	"os"
	"runtime"

	"github.com/gobuffalo/clara/genny/helpers"
	"github.com/gobuffalo/meta"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/syncx"
	"github.com/pkg/errors"
)

type Options struct {
	App         meta.App
	Versions    syncx.StringMap
	Out         io.Writer
	SkipBuffalo bool
	SkipNode    bool
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	if opts.App.IsZero() {
		opts.App = meta.New(".")
	}
	if opts.Out == nil {
		opts.Out = os.Stdout
	}
	if _, ok := opts.Versions.Load("go"); !ok {
		opts.Versions.Store("go", runtime.Version())
	}
	return nil
}

func (opts *Options) render(s string, ctx *plush.Context) error {
	s, err := templates.FindString(s)
	if err != nil {
		return errors.WithStack(err)
	}
	ctx.Set("opts", opts)
	return helpers.Render(opts.Out, s, ctx)
}
