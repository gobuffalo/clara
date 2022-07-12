package rx

import (
	"os"

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
