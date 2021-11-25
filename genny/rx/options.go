package rx

import (
	"os"
	"runtime"

	"github.com/gobuffalo/meta/v2"
	"github.com/gobuffalo/plush/v4"
)

type Options struct {
	App         *meta.App
	Versions    StringMap
	Out         Writer
	SkipBuffalo bool
	SkipNode    bool
	SkipDB      bool
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	if opts.App == nil || opts.App.Info.IsZero() {
		var err error
		opts.App, err = meta.NewDir(".")
		if err != nil {
			return err
		}
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
	s, err := templateFeeder(s)
	if err != nil {
		return err
	}
	ctx.Set("opts", opts)
	ctx.Set("partialFeeder", templateFeeder)
	return opts.Out.Render(s, ctx)
}
