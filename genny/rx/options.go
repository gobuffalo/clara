package rx

import (
	"io"
	"os"
	"runtime"

	"github.com/gobuffalo/doctor/genny/helpers"
	"github.com/gobuffalo/meta"
	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

type Options struct {
	App       meta.App
	GoVersion string
	Out       io.Writer
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	if opts.App.IsZero() {
		opts.App = meta.New(".")
	}
	if opts.Out == nil {
		opts.Out = os.Stdout
	}
	if len(opts.GoVersion) == 0 {
		opts.GoVersion = runtime.Version()
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
