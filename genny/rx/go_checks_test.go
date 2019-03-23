package rx

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/gobuffalo/clara/genny/helpers"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/genny/gentest"
	"github.com/gobuffalo/meta"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func Test_goPathCheck_Mods(t *testing.T) {
	r := require.New(t)
	envy.Temp(func() {
		envy.Set(envy.GO111MODULE, "on")

		run := gentest.NewRunner()

		bb := &bytes.Buffer{}
		run.WithRun(goPathCheck(&Options{
			Out: bb,
		}))
		r.NoError(run.Run())
		r.Contains(bb.String(), "You are using Go Modules")
	})
}

func Test_goPathCheck_Valid(t *testing.T) {
	r := require.New(t)
	envy.Temp(func() {
		envy.Set(envy.GO111MODULE, "off")

		run := gentest.NewRunner()

		bb := &bytes.Buffer{}

		envy.Set("GOPATH", "/foo")
		app := meta.New(".")
		app.Pwd = "/foo/src/bar"
		run.WithRun(goPathCheck(&Options{
			App: app,
			Out: bb,
		}))
		r.NoError(run.Run())
		r.Contains(bb.String(), "operating inside of your GOPATH")
	})
}

func Test_goPathCheck_Invalid(t *testing.T) {
	r := require.New(t)
	envy.Temp(func() {
		envy.Set(envy.GO111MODULE, "off")

		run := gentest.NewRunner()

		bb := &bytes.Buffer{}

		envy.Set("GOPATH", "/foo")
		app := meta.New(".")
		app.Pwd = "asdfasdf"
		run.WithRun(goPathCheck(&Options{
			App: app,
			Out: bb,
		}))
		r.NoError(run.Run())
		r.Contains(bb.String(), "Things to check")
	})
}

func Test_pkg_Mods(t *testing.T) {
	r := require.New(t)
	envy.Temp(func() {
		envy.Set(envy.GO111MODULE, "on")

		bb := &bytes.Buffer{}
		run := gentest.NewRunner()
		run.WithRun(goPkgCheck(&Options{
			Out: bb,
		}))

		r.NoError(run.Run())

		r.Contains(bb.String(), "You are using Go Modules")
	})
}

func Test_pkg_Dep(t *testing.T) {
	r := require.New(t)
	envy.Temp(func() {
		envy.Set(envy.GO111MODULE, "off")

		bb := &bytes.Buffer{}
		run := gentest.NewRunner()
		app := meta.New(".")
		app.WithDep = true
		run.LookPathFn = func(s string) (string, error) {
			return s, nil
		}
		run.WithRun(goPkgCheck(&Options{
			App: app,
			Out: bb,
		}))

		r.NoError(run.Run())

		r.Contains(bb.String(), "You are using Dep")
	})
}

func Test_pkg_Dep_notFound(t *testing.T) {
	r := require.New(t)
	envy.Temp(func() {
		envy.Set(envy.GO111MODULE, "off")

		bb := &bytes.Buffer{}
		run := gentest.NewRunner()
		app := meta.New(".")
		app.WithDep = true
		run.LookPathFn = func(s string) (string, error) {
			return s, errors.New("oops")
		}
		run.WithRun(goPkgCheck(&Options{
			App: app,
			Out: bb,
		}))

		r.NoError(run.Run())

		r.Contains(bb.String(), "`dep` executable could not be found")
	})
}

func Test_goPathBinCheck_Valid(t *testing.T) {
	r := require.New(t)
	envy.Temp(func() {
		envy.Set("PATH", filepath.Join("foo", "bin"))
		envy.Set("GOPATH", "foo")

		run := gentest.NewRunner()

		bb := &bytes.Buffer{}
		run.WithRun(goPathBinCheck(&Options{
			Out: bb,
		}))
		r.NoError(run.Run())

		r.Contains(bb.String(), helpers.SUCCESS)
	})
}

func Test_goPathBinCheck_Invalid(t *testing.T) {
	r := require.New(t)
	envy.Temp(func() {
		envy.Set("GOPATH", "/foo")

		run := gentest.NewRunner()

		bb := &bytes.Buffer{}
		run.WithRun(goPathBinCheck(&Options{
			Out: bb,
		}))
		r.NoError(run.Run())

		r.Contains(bb.String(), helpers.ERROR)
	})
}
