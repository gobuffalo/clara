package rx

import (
	"bytes"
	"testing"

	"github.com/gobuffalo/genny/gentest"
	"github.com/gobuffalo/syncx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func Test_buffaloChecks_Success(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := syncx.StringMap{}
	v.Store("buffalo", "1.0.0")
	run.With(buffaloChecks(&Options{
		Out:      bb,
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `buffalo` executable was found")
	r.Contains(res, "Your version of Buffalo, 1.0.0, meets the minimum requirements.")
}

func Test_buffaloChecks_Failure(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := syncx.StringMap{}
	v.Store("buffalo", "0.0.0")
	run.With(buffaloChecks(&Options{
		Out:      bb,
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, errors.New("oops")
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `buffalo` executable could not be found")
	r.Contains(res, "Your version of Buffalo, 0.0.0, does not meet the minimum requirements.")
}
