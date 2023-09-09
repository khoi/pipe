package manifest

import "context"

type Manifest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Pipe        Pipe     `json:"pipe"`
	Output      Output   `json:"output"`
	Tags        []string `json:"tags"`
}

func (m Manifest) Execute(ctx context.Context, input *string) error {
	cmd := m.Pipe.Command(ctx, input)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return m.Output.Write(ctx, string(output))
}
