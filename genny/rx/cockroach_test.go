package rx

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"

	"github.com/gobuffalo/genny/gentest"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func Test_cockroachChecks_Success(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := StringMap{}
	run.ExecFn = func(c *exec.Cmd) error {
		a := strings.Join(c.Args, " ")
		if a != "cockroach version" {
			return nil
		}
		c.Stdout.Write([]byte(crv))
		return nil
	}
	run.With(cockroachChecks(&Options{
		Out:      NewWriter(bb),
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, nil
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `cockroach` executable was found")
	r.Contains(res, "Your version of Cockroach, 2.0.5, meets the minimum requirements.")
}

func Test_cockroachChecks_Failure(t *testing.T) {
	r := require.New(t)

	run := gentest.NewRunner()
	bb := &bytes.Buffer{}

	v := StringMap{}
	v.Store("cockroach", "0.0.0")
	run.With(cockroachChecks(&Options{
		Out:      NewWriter(bb),
		Versions: v,
	}))

	run.LookPathFn = func(s string) (string, error) {
		return s, errors.New("oops")
	}

	r.NoError(run.Run())

	res := bb.String()
	r.Contains(res, "The `cockroach` executable could not be found")
}

const crv = `Build Tag:    v2.0.5
Build Time:   2018/08/13 20:55:02
Distribution: CCL
Platform:     darwin amd64 (x86_64-apple-darwin17.7.0)
Go Version:   go1.10.3
C Compiler:   4.2.1 Compatible Apple LLVM 9.1.0 (clang-902.0.39.2)
Build SHA-1:  3f2f0f5eea9b9d552e471eba8c37504a0595342f
Build Type:   development`
