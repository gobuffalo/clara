package rx

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/gobuffalo/genny"
)

func nodeChecks(opts *Options) *genny.Generator {
	t := Tool{
		Name:    "Node",
		Bin:     "node",
		Minimum: []string{">=1.11"},
		Partial: "node/_help.plush",
		Version: func() (string, error) {
			if v, ok := opts.Versions.Load("node"); ok {
				return v, nil
			}
			bb := &bytes.Buffer{}
			c := exec.Command("node", "--version")
			c.Stdout = bb
			c.Stderr = bb
			if err := c.Run(); err != nil {
				return "", err
			}
			v := strings.TrimSpace(bb.String())
			return v, nil
		},
	}

	g := t.Generator(opts)
	return g
}

func yarnChecks(opts *Options) *genny.Generator {
	t := Tool{
		Name:    "Yarn",
		Bin:     "yarnpkg",
		Minimum: []string{">=1.12"},
		Partial: "node/_yarn.plush",
		Version: func() (string, error) {
			if v, ok := opts.Versions.Load("yarn"); ok {
				return v, nil
			}
			bb := &bytes.Buffer{}
			c := exec.Command("yarnpkg", "--version")
			c.Stdout = bb
			c.Stderr = bb
			if err := c.Run(); err != nil {
				return "", err
			}
			v := strings.TrimSpace(bb.String())
			return v, nil
		},
	}

	g := t.Generator(opts)
	return g
}

func npmChecks(opts *Options) *genny.Generator {
	t := Tool{
		Name:    "NPM",
		Bin:     "npm",
		Minimum: []string{">=6.0.0", ">=7.0.0"},
		Partial: "node/_npm.plush",
		Version: func() (string, error) {
			if v, ok := opts.Versions.Load("npm"); ok {
				return v, nil
			}
			bb := &bytes.Buffer{}
			c := exec.Command("npm", "--version")
			c.Stdout = bb
			c.Stderr = bb
			if err := c.Run(); err != nil {
				return "", err
			}
			v := strings.TrimSpace(bb.String())
			return v, nil
		},
	}

	g := t.Generator(opts)
	return g
}
