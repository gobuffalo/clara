package rx

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"

	"github.com/gobuffalo/genny/gentest"
	"github.com/gobuffalo/syncx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func Test_sqliteChecks_Success(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := syncx.StringMap{}
	run.ExecFn = func(c *exec.Cmd) error {
		a := strings.Join(c.Args, " ")
		if a != "sqlite3 --version" {
			return nil
		}
		c.Stdout.Write([]byte("3.24.0 2018-06-04 14:10:15 95fbac39baaab1c3a84fdfc82ccb7f42398b2e92f18a2a57bce1d4a713cbaapl"))
		return nil
	}
	run.With(sqliteChecks(&Options{
		Out:      NewWriter(bb),
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `sqlite3` executable was found")
	r.Contains(res, "Your version of SQLite3, 3.24.0, meets the minimum requirements.")
}

func Test_sqliteChecks_Failure(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := syncx.StringMap{}
	v.Store("sqlite", "0.0.0")
	run.With(sqliteChecks(&Options{
		Out:      NewWriter(bb),
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, errors.New("oops")
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `sqlite3` executable could not be found")
}
