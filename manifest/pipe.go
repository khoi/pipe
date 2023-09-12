package manifest

import (
	"context"
	"io"
	"strings"

	"github.com/khoi/pipe/funk"
)

type Pipe struct {
	Exec  string     `json:"exec"`
	Args  []Argument `json:"args"`
	Stdin *Argument  `json:"stdin,omitempty"`
}

func (p Pipe) Command(ctx context.Context, input *string) *Cmd {
	args := funk.FlatMap(p.Args, func(a Argument) []string {
		return a.Value(input)
	})

	cmd := newCommand(ctx, "bash", "-c", p.Exec+" "+strings.Join(args, " "))
	if p.Stdin != nil {
		cmd.Stdin = &inputReader{input: p.Stdin.Value(input)}
	}
	return cmd
}

type inputReader struct {
	input []string
	idx   int
}

func (r *inputReader) Read(p []byte) (int, error) {
	if r.idx >= len(r.input) {
		return 0, io.EOF
	}
	n := copy(p, r.input[r.idx])
	r.idx++
	return n, nil
}
