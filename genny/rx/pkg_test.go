package rx

import (
	"bytes"
	"testing"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/genny/gentest"
	"github.com/gobuffalo/meta"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func Test_pkg_Mods(t *testing.T) {
	r := require.New(t)
	envy.Temp(func() {
		envy.Set(envy.GO111MODULE, "on")

		bb := &bytes.Buffer{}
		run := gentest.NewRunner()
		run.WithRun(pkgManagement(&Options{
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
		run.WithRun(pkgManagement(&Options{
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
		run.WithRun(pkgManagement(&Options{
			App: app,
			Out: bb,
		}))

		r.NoError(run.Run())

		r.Contains(bb.String(), "`dep` executable could not be found")
	})
}
