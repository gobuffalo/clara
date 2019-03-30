package rx

import (
	"bytes"
	"testing"

	"github.com/gobuffalo/genny/gentest"
	"github.com/gobuffalo/syncx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func Test_nodeChecks_Success(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := syncx.StringMap{}
	v.Store("node", "1.12.0")
	run.With(nodeChecks(&Options{
		Out:      NewWriter(bb),
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `node` executable was found")
	r.Contains(res, "Your version of Node, 1.12.0, meets the minimum requirements.")
}

func Test_nodeChecks_Failure(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := syncx.StringMap{}
	v.Store("node", "0.0.0")
	run.With(nodeChecks(&Options{
		Out:      NewWriter(bb),
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, errors.New("oops")
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `node` executable could not be found")
}
