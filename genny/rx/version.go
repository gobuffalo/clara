package rx

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/gobuffalo/genny/v2"
)

func cmdVersion(r *genny.Runner, bin string, args ...string) (string, error) {
	bb := &bytes.Buffer{}
	c := exec.Command(bin, args...)
	c.Stdout = bb
	c.Stderr = bb
	if err := r.Exec(c); err != nil {
		return "", err
	}
	v := strings.TrimSpace(bb.String())
	return v, nil
}
