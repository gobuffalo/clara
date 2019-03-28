package rx

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Writer_Tabs(t *testing.T) {
	r := require.New(t)

	lines := [][]string{
		{"a", "b"},
		{"c", "d"},
	}

	bb := &bytes.Buffer{}

	w := NewWriter(bb)
	r.NoError(w.Tabs(lines))

	res := bb.String()
	r.Equal("a b\nc d\n", res)
}
