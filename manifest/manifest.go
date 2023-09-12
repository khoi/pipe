package manifest

import (
	"context"
	"encoding/json"
	"io/fs"
)

type Manifest struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Pipe        Pipe     `json:"pipe"`
	Output      Output   `json:"output"`
	Tags        []string `json:"tags"`
	filesystem  fs.FS    `json:"-"`
}

func NewFromBytes(filesystem fs.FS, id string, bytes []byte) (Manifest, error) {
	var m Manifest
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return m, err
	}
	m.ID = id
	m.filesystem = filesystem
	return m, nil
}

func (m Manifest) Execute(ctx context.Context, input *string) (string, error) {
	cmd := m.Pipe.Command(ctx, m.filesystem, input)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	out := string(output)
	return out, m.Output.Write(ctx, out)
}
