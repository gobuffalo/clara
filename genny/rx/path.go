package rx

import (
	"path/filepath"
	"strings"

	"github.com/gobuffalo/clara/genny/helpers"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/genny"
)

func checkPath(opts *Options) genny.RunFn {
	return func(r *genny.Runner) error {
		helpers.Header(opts.Out, "Checking PATH")
		path := envy.Get("PATH", "")

		ctx := helpers.Context(opts)
		if strings.Contains(path, filepath.Join(envy.GoPath(), "bin")) {
			return opts.render("valid_path.plush", ctx)
		}
		return opts.render("invalid_path.plush", ctx)
	}
}
