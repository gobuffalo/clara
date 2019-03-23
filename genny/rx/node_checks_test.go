package rx

import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/gobuffalo/clara/genny/helpers"
	"github.com/gobuffalo/genny/gentest"
	"github.com/gobuffalo/meta"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func Test_nodeChecks_InApp_Missing(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.InApp = true

	bb := &bytes.Buffer{}
	run.With(nodeChecks(&Options{
		Out: bb,
		App: app,
	}))
	run.LookPathFn = func(s string) (string, error) {
		return s, errors.New("missing")
	}

	r.NoError(run.Run())
	r.Contains(bb.String(), helpers.ERROR)
}

func Test_nodeChecks_InApp_Invalid_Version(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.InApp = true

	bb := &bytes.Buffer{}
	opts := &Options{
		Out: bb,
		App: app,
	}
	opts.Versions.Store("node", "0.0.1")
	run.With(nodeChecks(opts))
	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}
	r.NoError(run.Run())
	r.Contains(bb.String(), helpers.ERROR)
}

func Test_nodeChecks_InApp_Valid_Version(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.InApp = true

	bb := &bytes.Buffer{}
	run.With(npmChecks(&Options{
		Out: bb,
		App: app,
	}))
	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}
	run.ExecFn = func(cmd *exec.Cmd) error {
		if cmd.Stdout != nil {
			cmd.Stdout.Write([]byte("v11.9.0"))
		}
		return nil
	}

	r.NoError(run.Run())
	r.Contains(bb.String(), helpers.SUCCESS)
}

func Test_npmChecks_InApp_Missing(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.InApp = true

	bb := &bytes.Buffer{}
	run.With(npmChecks(&Options{
		Out: bb,
		App: app,
	}))
	run.LookPathFn = func(s string) (string, error) {
		return s, errors.New("missing")
	}

	r.NoError(run.Run())
	r.Contains(bb.String(), helpers.ERROR)
}

func Test_npmChecks_InApp_Invalid_Version(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.InApp = true

	bb := &bytes.Buffer{}
	opts := &Options{
		Out: bb,
		App: app,
	}
	opts.Versions.Store("npm", "0.0.1")
	run.With(npmChecks(opts))
	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}
	r.NoError(run.Run())
	r.Contains(bb.String(), helpers.ERROR)
}

func Test_npmChecks_InApp_Valid_Version(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.InApp = true

	bb := &bytes.Buffer{}
	opts := &Options{
		Out: bb,
		App: app,
	}
	opts.Versions.Store("npm", "v6.5.0")
	run.With(npmChecks(opts))
	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}
	r.NoError(run.Run())
	r.Contains(bb.String(), helpers.SUCCESS)
}

func Test_yarnChecks_InApp_Missing(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.WithYarn = true
	app.InApp = true

	bb := &bytes.Buffer{}
	run.With(yarnChecks(&Options{
		Out: bb,
		App: app,
	}))
	run.LookPathFn = func(s string) (string, error) {
		return s, errors.New("missing")
	}

	r.NoError(run.Run())
	r.Contains(bb.String(), helpers.ERROR)
}

func Test_yarnChecks_InApp_Invalid_Version(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.WithYarn = true
	app.InApp = true

	bb := &bytes.Buffer{}
	run.With(yarnChecks(&Options{
		Out: bb,
		App: app,
	}))
	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}
	run.ExecFn = func(cmd *exec.Cmd) error {
		if cmd.Stdout != nil {
			cmd.Stdout.Write([]byte("v1.1.0"))
		}
		return nil
	}

	r.NoError(run.Run())
	r.Contains(bb.String(), helpers.ERROR)
}

func Test_yarnChecks_InApp_Valid_Version(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.WithYarn = true
	app.InApp = true

	bb := &bytes.Buffer{}
	run.With(yarnChecks(&Options{
		Out: bb,
		App: app,
	}))
	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}
	run.ExecFn = func(cmd *exec.Cmd) error {
		if cmd.Stdout != nil {
			cmd.Stdout.Write([]byte("v1.12.0"))
		}
		return nil
	}

	r.NoError(run.Run())
	r.Contains(bb.String(), helpers.SUCCESS)
}
