package rx

import (
	"bytes"
	"testing"

	"github.com/gobuffalo/genny/v2/gentest"
	"github.com/stretchr/testify/require"
)

func Test_pkg_Mods(t *testing.T) {
	r := require.New(t)
	bb := &bytes.Buffer{}
	run := gentest.NewRunner()
	run.WithRun(goPkgCheck(&Options{
		Out: NewWriter(bb),
	}))

	r.NoError(run.Run())

	r.Contains(bb.String(), "You are using Go Modules")
}
