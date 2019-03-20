package rx

import (
	"bytes"
	"testing"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/genny/gentest"
	"github.com/gobuffalo/meta"
	"github.com/stretchr/testify/require"
)

func Test_gopath_Mods(t *testing.T) {
	r := require.New(t)
	envy.Temp(func() {
		envy.Set(envy.GO111MODULE, "on")

		run := gentest.NewRunner()

		bb := &bytes.Buffer{}
		run.WithRun(gopath(&Options{
			Out: bb,
		}))
		r.NoError(run.Run())
		r.Contains(bb.String(), "You are using Go Modules")
	})
}

func Test_gopath_Valid(t *testing.T) {
	r := require.New(t)
	envy.Temp(func() {
		envy.Set(envy.GO111MODULE, "off")

		run := gentest.NewRunner()

		bb := &bytes.Buffer{}

		envy.Set("GOPATH", "/foo")
		app := meta.New(".")
		app.Pwd = "/foo/src/bar"
		run.WithRun(gopath(&Options{
			App: app,
			Out: bb,
		}))
		r.NoError(run.Run())
		r.Contains(bb.String(), "operating inside of your GOPATH")
	})
}

func Test_gopath_Invalid(t *testing.T) {
	r := require.New(t)
	envy.Temp(func() {
		envy.Set(envy.GO111MODULE, "off")

		run := gentest.NewRunner()

		bb := &bytes.Buffer{}

		envy.Set("GOPATH", "/foo")
		app := meta.New(".")
		app.Pwd = "asdfasdf"
		run.WithRun(gopath(&Options{
			App: app,
			Out: bb,
		}))
		r.NoError(run.Run())
		r.Contains(bb.String(), "Things to check")
	})
}
