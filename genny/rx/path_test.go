package rx

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/gobuffalo/clara/genny/helpers"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/genny/gentest"
	"github.com/stretchr/testify/require"
)

func Test_checkPath_Valid(t *testing.T) {
	r := require.New(t)
	envy.Temp(func() {
		envy.Set("PATH", filepath.Join("foo", "bin"))
		envy.Set("GOPATH", "foo")

		run := gentest.NewRunner()

		bb := &bytes.Buffer{}
		run.WithRun(checkPath(&Options{
			Out: bb,
		}))
		r.NoError(run.Run())

		r.Contains(bb.String(), helpers.SUCCESS)
	})
}

func Test_checkPath_Invalid(t *testing.T) {
	r := require.New(t)
	envy.Temp(func() {
		envy.Set("GOPATH", "/foo")

		run := gentest.NewRunner()

		bb := &bytes.Buffer{}
		run.WithRun(checkPath(&Options{
			Out: bb,
		}))
		r.NoError(run.Run())

		r.Contains(bb.String(), helpers.ERROR)
	})
}
