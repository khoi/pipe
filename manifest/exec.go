package manifest

import (
	"context"
	"os"
	"os/exec"
)

// Cmd is a thin abstraction around os/exec.Cmd that sends os.Interrupt when
// the context is canceled (instead of os.Kill in the stdlib).
type Cmd struct {
	*exec.Cmd
	ctx context.Context
}

func newCommand(ctx context.Context, name string, arg ...string) *Cmd {
	return &Cmd{
		Cmd: exec.Command(name, arg...),
		ctx: ctx,
	}
}

func (c *Cmd) String() string {
	return c.Cmd.String()
}

func (c *Cmd) CombinedOutput() ([]byte, error) {
	go interruptOnCancel(c)
	return c.Cmd.CombinedOutput()
}

func (c *Cmd) Output() ([]byte, error) {
	go interruptOnCancel(c)
	return c.Cmd.Output()
}

func (c *Cmd) Run() error {
	go interruptOnCancel(c)
	return c.Cmd.Run()
}

func (c *Cmd) Start() error {
	go interruptOnCancel(c)
	return c.Cmd.Start()
}

func (c *Cmd) Wait() error {
	go interruptOnCancel(c)
	return c.Cmd.Wait()
}

func interruptOnCancel(c *Cmd) {
	if c.ctx == nil {
		return
	}
	<-c.ctx.Done()

	if c.Cmd.Process == nil {
		return
	}
	_ = c.Cmd.Process.Signal(os.Interrupt)
}
