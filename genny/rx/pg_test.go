package rx

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/gobuffalo/genny/gentest"
	"github.com/stretchr/testify/require"
)

func Test_postgresChecks_Success(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := StringMap{}
	run.ExecFn = func(c *exec.Cmd) error {
		a := strings.Join(c.Args, " ")
		if a != "postgres --version" {
			return nil
		}
		c.Stdout.Write([]byte("postgres (PostgreSQL) 10.5"))
		return nil
	}
	run.With(postgresChecks(&Options{
		Out:      NewWriter(bb),
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `postgres` executable was found")
	r.Contains(res, "Your version of PostgreSQL, 10.5.0, meets the minimum requirements.")
}

func Test_postgresChecks_Failure(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := StringMap{}
	v.Store("postgres", "0.0.0")
	run.With(postgresChecks(&Options{
		Out:      NewWriter(bb),
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, fmt.Errorf("oops")
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `postgres` executable could not be found")
}
