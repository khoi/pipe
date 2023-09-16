package manifest

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

func uuidGenerate(_ context.Context, input *string) (string, error) {
	id := uuid.New()
	return id.String(), nil
}

var UUIDGenerate = Manifest{
	ID:          "uuid-generate",
	Name:        "UUID Generate",
	Description: "Generate a UUID",
	Pipe: Pipe{
		Handler: uuidGenerate,
	},
	Output: Replace,
	Tags:   []string{"uuid", "generate", "uuid-generate"},
}
