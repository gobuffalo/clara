package rx

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/gobuffalo/plush"
)

func (w Writer) RenderE(err error) error {
	s := color.RedString(fmt.Sprintf("%s %v", ERROR, err))
	s = strings.TrimSpace(s)
	s += "\n\n"
	return w.WriteString(s)
}

func (w Writer) Render(s string, ctx *plush.Context) error {
	ctx.Set("partialFeeder", templates.FindString)
	s, err := plush.Render(s, ctx)
	if err != nil {
		return err
	}
	s = strings.TrimSpace(s)
	s += "\n\n"
	return w.WriteString(s)
}
