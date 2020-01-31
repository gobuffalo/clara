package rx

import (
	"os"
	"runtime"

	"github.com/gobuffalo/meta"
	"github.com/gobuffalo/plush/v4"
)

type Options struct {
	App         meta.App
	Versions    StringMap
	Out         Writer
	SkipBuffalo bool
	SkipNode    bool
	SkipDB      bool
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	if opts.App.IsZero() {
		opts.App = meta.New(".")
	}
	if opts.Out.Writer == nil {
		opts.Out = NewWriter(os.Stdout)
	}
	if _, ok := opts.Versions.Load("go"); !ok {
		opts.Versions.Store("go", runtime.Version())
	}
	return nil
}

func (opts *Options) render(s string, ctx *plush.Context) error {
	s, err := templates.FindString(s)
	if err != nil {
		return err
	}
	ctx.Set("opts", opts)
	ctx.Set("partialFeeder", templates.FindString)
	return opts.Out.Render(s, ctx)
}
