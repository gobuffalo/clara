package rx

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gobuffalo/genny/gentest"
	"github.com/stretchr/testify/require"
)

func Test_npmChecks_Success(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := StringMap{}
	v.Store("npm", "7.0.0")
	run.With(npmChecks(&Options{
		Out:      NewWriter(bb),
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `npm` executable was found")
	r.Contains(res, "Your version of NPM, 7.0.0, meets the minimum requirements.")
}

func Test_npmChecks_Failure(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := StringMap{}
	v.Store("npm", "0.0.0")
	run.With(npmChecks(&Options{
		Out:      NewWriter(bb),
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, fmt.Errorf("oops")
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `npm` executable could not be found")
}
