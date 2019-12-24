package rx

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/gobuffalo/genny/v2/gentest"
	"github.com/stretchr/testify/require"
)

func Test_mysqlChecks_Success(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := StringMap{}
	run.ExecFn = func(c *exec.Cmd) error {
		a := strings.Join(c.Args, " ")
		if a != "mysql --version" {
			return nil
		}
		c.Stdout.Write([]byte("mysql  Ver 8.0.12 for osx10.13 on x86_64 (Homebrew)"))
		return nil
	}
	run.With(mysqlChecks(&Options{
		Out:      NewWriter(bb),
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `mysql` executable was found")
	r.Contains(res, "Your version of MySQL, 8.0.12, meets the minimum requirements.")
}

func Test_mysqlChecks_Failure(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := StringMap{}
	v.Store("mysql", "0.0.0")
	run.With(mysqlChecks(&Options{
		Out:      NewWriter(bb),
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, fmt.Errorf("oops")
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `mysql` executable could not be found")
}
