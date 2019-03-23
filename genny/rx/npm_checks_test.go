package rx

import (
	"bytes"
	"testing"

	"github.com/gobuffalo/genny/gentest"
	"github.com/gobuffalo/syncx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func Test_npmChecks_Success(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := syncx.StringMap{}
	v.Store("npm", "7.0.0")
	run.With(npmChecks(&Options{
		Out:      bb,
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

	v := syncx.StringMap{}
	v.Store("npm", "0.0.0")
	run.With(npmChecks(&Options{
		Out:      bb,
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, errors.New("oops")
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `npm` executable could not be found")
	r.Contains(res, "Your version of NPM, 0.0.0, does not meet the minimum requirements.")
}
