package manifest

import (
	"context"
	"io"
	"io/fs"
	"os"
	"os/exec"
)

// Cmd is a thin abstraction around os/exec.Cmd that sends os.Interrupt when
// the context is canceled (instead of os.Kill in the stdlib).
type Cmd struct {
	*exec.Cmd
	ctx context.Context
}

func newCommand(ctx context.Context, execDir fs.FS, name string, arg ...string) (*Cmd, func()) {
	cmd := &Cmd{
		Cmd: exec.Command(name, arg...),
		ctx: ctx,
	}
	cmd.Cmd.Env = os.Environ()
	cleanup := ensureExecDir(cmd, execDir)
	return cmd, cleanup
}

func ensureExecDir(cmd *Cmd, execDir fs.FS) func() {
	tempDir, err := os.MkdirTemp("", "pipe_exec")
	if err != nil {
		panic(err)
	}
	cleanup := func() {
		_ = os.RemoveAll(tempDir)
	}
	cmd.Dir = tempDir

	fs.WalkDir(execDir, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		f, err := execDir.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		out, err := os.Create(tempDir + "/" + path)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, f)
		return err
	})

	return cleanup
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
