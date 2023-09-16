package manifest

import (
	"context"
	"io/fs"
	"strings"

	"github.com/khoi/pipe/funk"
)

type Pipe struct {
	Exec  string     `json:"exec"`
	Args  []Argument `json:"args"`
	Stdin *Argument  `json:"stdin,omitempty"`
}

func (p Pipe) Command(ctx context.Context, execDir fs.FS, input *string) (*Cmd, func()) {
	args := funk.FlatMap(p.Args, func(a Argument) []string {
		return a.Value(input)
	})

	cmd, cleanup := newCommand(ctx, execDir, "bash", "-c", p.Exec+" "+strings.Join(args, " "))
	if p.Stdin != nil {
		in := strings.Join(p.Stdin.Value(input), "\n")
		cmd.Stdin = strings.NewReader(in)
	}
	return cmd, cleanup
}
