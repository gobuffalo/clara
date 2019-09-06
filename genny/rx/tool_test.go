package rx

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Tool_AcceptVersion(t *testing.T) {
	table := []struct {
		mins    []string
		version string
		Error   bool
		Pass    bool
	}{
		{mins: []string{">= 1.10.5", ">= 1.11.5", ">= 1.12"}, version: "1.12.1", Pass: true},
		{mins: []string{">= v1.10.5", ">= v1.11.5", ">= v1.12"}, version: "1.12.1", Pass: true},
		{mins: []string{"1.10.5", "1.11.5", ">= 1.12"}, version: "1.12.1", Pass: true},
		{mins: []string{"1.10.5", "1.11.5", "1.12"}, version: "1.13.1", Pass: false},
		{mins: []string{"1.10.5", "1.11.5", "1.12"}, version: "1.2.1", Pass: false},
	}

	for _, tt := range table {
		t.Run(strings.Join(tt.mins, "-")+tt.version, func(st *testing.T) {
			r := require.New(st)
			to := Tool{
				Minimum: tt.mins,
			}
			b, err := to.AcceptVersion(tt.version)
			if tt.Error {
				r.Error(err)
			} else {
				r.NoError(err)
			}
			if tt.Pass {
				r.True(b)
			} else {
				r.False(b)
			}
		})
	}
}
