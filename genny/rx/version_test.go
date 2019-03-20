package rx

import (
	"bytes"
	"testing"

	"github.com/gobuffalo/genny/gentest"
	"github.com/stretchr/testify/require"
)

func Test_goVersion(t *testing.T) {
	table := []string{"1.10.8", "1.11.4", "1.12", "1.12.1"}

	for _, tt := range table {
		t.Run(tt, func(st *testing.T) {
			r := require.New(st)
			bb := &bytes.Buffer{}

			run := gentest.NewRunner()
			run.WithRun(goVersion(&Options{
				GoVersion: tt,
				Out:       bb,
			}))

			r.NoError(run.Run())
			r.NotContains(bb.String(), "does not meet the minimum")
		})
	}
}

func Test_goVersion_Below_Min(t *testing.T) {
	table := []string{"1.9.7", "1.8", "1.10", "1.11.2"}

	for _, tt := range table {
		t.Run(tt, func(st *testing.T) {
			r := require.New(st)
			bb := &bytes.Buffer{}

			run := gentest.NewRunner()
			run.WithRun(goVersion(&Options{
				GoVersion: tt,
				Out:       bb,
			}))

			r.NoError(run.Run())
			r.Contains(bb.String(), "does not meet the minimum")
		})
	}
}
