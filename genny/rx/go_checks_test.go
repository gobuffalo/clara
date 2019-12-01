package rx

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/genny/gentest"
	"github.com/stretchr/testify/require"
)

func Test_pkg_Mods(t *testing.T) {
	r := require.New(t)
	envy.Temp(func() {
		envy.Set(envy.GO111MODULE, "on")

		bb := &bytes.Buffer{}
		run := gentest.NewRunner()
		run.WithRun(goPkgCheck(&Options{
			Out: NewWriter(bb),
		}))

		r.NoError(run.Run())

		r.Contains(bb.String(), "You are using Go Modules")
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
			Out: NewWriter(bb),
		}))
		r.NoError(run.Run())

		r.Contains(bb.String(), SUCCESS)
	})
}

func Test_goPathBinCheck_Invalid(t *testing.T) {
	r := require.New(t)
	envy.Temp(func() {
		envy.Set("GOPATH", "/foo")

		run := gentest.NewRunner()

		bb := &bytes.Buffer{}
		run.WithRun(goPathBinCheck(&Options{
			Out: NewWriter(bb),
		}))
		r.NoError(run.Run())

		r.Contains(bb.String(), ERROR)
	})
}
