package rx

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gobuffalo/genny/gentest"
	"github.com/stretchr/testify/require"
)

func Test_yarnChecks_Success(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := StringMap{}
	v.Store("yarn", "1.15.0")
	run.With(yarnChecks(&Options{
		Out:      NewWriter(bb),
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `yarnpkg` executable was found")
	r.Contains(res, "Your version of Yarn, 1.15.0, meets the minimum requirements.")
}

func Test_yarnChecks_Failure(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := StringMap{}
	v.Store("yarn", "0.0.0")
	run.With(yarnChecks(&Options{
		Out:      NewWriter(bb),
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, fmt.Errorf("oops")
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `yarnpkg` executable could not be found")
}
