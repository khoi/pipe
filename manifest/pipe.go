package manifest

import (
	"context"
	"io"

	"github.com/khoi/pipe/funk"
)

type Pipe struct {
	Exec  string     `json:"exec"`
	Args  []Argument `json:"args"`
	Stdin *string    `json:"stdin,omitempty"`
}

func (p Pipe) Command(ctx context.Context, input *string) *Cmd {
	args := funk.FlatMap(p.Args, func(a Argument) []string {
		return a.Value(input)
	})

	cmd := newCommand(ctx, p.Exec, args...)
	if p.Stdin != nil {
		cmd.Stdin = &inputReader{input: *p.Stdin}
	}
	return cmd
}

type inputReader struct {
	input string
	idx   int
}

func (r *inputReader) Read(p []byte) (int, error) {
	if r.idx >= len(r.input) {
		return 0, io.EOF
	}
	n := copy(p, r.input[r.idx:])
	r.idx += n
	return n, nil
}
