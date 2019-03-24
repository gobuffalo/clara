package rx

import (
	"fmt"
	"io"
	"strings"
)

func Header(w io.Writer, s string) {
	s = strings.TrimSpace(s)
	s = fmt.Sprintf("-> %s\n", s)
	w.Write([]byte(s))
}
