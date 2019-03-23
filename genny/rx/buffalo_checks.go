package rx

import (
	"bytes"
	"os/exec"
	"regexp"
	"strings"

	"github.com/gobuffalo/genny"
)

var bvrx = regexp.MustCompile(`v\d+\.\d+\.\d+`)

func buffaloChecks(opts *Options) *genny.Generator {
	t := Tool{
		Name:    "Buffalo",
		Bin:     "buffalo",
		Minimum: []string{">=0.14.2"},
		Partial: "go/_help.plush",
		Version: func() (string, error) {
			if v, ok := opts.Versions.Load("buffalo"); ok {
				return v, nil
			}
			bb := &bytes.Buffer{}
			c := exec.Command("buffalo", "version")
			c.Stdout = bb
			c.Stderr = bb
			if err := c.Run(); err != nil {
				return "", err
			}
			v := bvrx.FindString(bb.String())
			v = strings.TrimSpace(v)
			return v, nil

		},
	}
	g := t.Generator(opts)
	return g
}
