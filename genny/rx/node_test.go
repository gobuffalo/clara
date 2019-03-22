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

func Test_checkNode_InApp_Missing(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.InApp = true

	bb := &bytes.Buffer{}
	run.WithRun(checkNode(&Options{
		Out: bb,
		App: app,
	}))
	run.LookPathFn = func(s string) (string, error) {
		return s, errors.New("missing")
	}

	r.NoError(run.Run())
	r.Contains(bb.String(), helpers.ERROR)
}

func Test_checkNode_InApp_Invalid_Version(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.InApp = true

	bb := &bytes.Buffer{}
	run.WithRun(checkNode(&Options{
		Out: bb,
		App: app,
	}))
	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}
	run.ExecFn = func(cmd *exec.Cmd) error {
		if cmd.Stdout != nil {
			cmd.Stdout.Write([]byte("v10.9.0"))
		}
		return nil
	}

	r.NoError(run.Run())
	r.Contains(bb.String(), helpers.ERROR)
}

func Test_checkNode_InApp_Valid_Version(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.InApp = true

	bb := &bytes.Buffer{}
	run.WithRun(checkNpm(&Options{
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

func Test_checkNpm_InApp_Missing(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.InApp = true

	bb := &bytes.Buffer{}
	run.WithRun(checkNpm(&Options{
		Out: bb,
		App: app,
	}))
	run.LookPathFn = func(s string) (string, error) {
		return s, errors.New("missing")
	}

	r.NoError(run.Run())
	r.Contains(bb.String(), helpers.ERROR)
}

func Test_checkNpm_InApp_Invalid_Version(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.InApp = true

	bb := &bytes.Buffer{}
	run.WithRun(checkNpm(&Options{
		Out: bb,
		App: app,
	}))
	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}
	run.ExecFn = func(cmd *exec.Cmd) error {
		if cmd.Stdout != nil {
			cmd.Stdout.Write([]byte("v6.5.0"))
		}
		return nil
	}

	r.NoError(run.Run())
	r.Contains(bb.String(), helpers.ERROR)
}

func Test_checkNpm_InApp_Valid_Version(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.InApp = true

	bb := &bytes.Buffer{}
	run.WithRun(checkNpm(&Options{
		Out: bb,
		App: app,
	}))
	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}
	run.ExecFn = func(cmd *exec.Cmd) error {
		if cmd.Stdout != nil {
			cmd.Stdout.Write([]byte("v6.7.0"))
		}
		return nil
	}

	r.NoError(run.Run())
	r.Contains(bb.String(), helpers.SUCCESS)
}

func Test_checkYarn_InApp_Missing(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.WithYarn = true
	app.InApp = true

	bb := &bytes.Buffer{}
	run.WithRun(checkYarn(&Options{
		Out: bb,
		App: app,
	}))
	run.LookPathFn = func(s string) (string, error) {
		return s, errors.New("missing")
	}

	r.NoError(run.Run())
	r.Contains(bb.String(), helpers.ERROR)
}

func Test_checkYarn_InApp_Invalid_Version(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.WithYarn = true
	app.InApp = true

	bb := &bytes.Buffer{}
	run.WithRun(checkYarn(&Options{
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

func Test_checkYarn_InApp_Valid_Version(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()

	app := meta.New(".")
	app.WithNodeJs = true
	app.WithYarn = true
	app.InApp = true

	bb := &bytes.Buffer{}
	run.WithRun(checkYarn(&Options{
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
