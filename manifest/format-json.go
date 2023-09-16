package manifest

import (
	"encoding/json"
	"log"

	"golang.org/x/net/context"
)

func formatJSON(_ context.Context, input *string) (string, error) {
	var data any
	if err := json.Unmarshal([]byte(*input), &data); err != nil {
		log.Println("Error unmarshalling json")
		return "", err
	}
	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("Error marshalling json")
		return "", err
	}
	return string(out), nil
}

var FormatJSON = Manifest{
	ID:          "format-json",
	Name:        "Format JSON",
	Description: "Format JSON document",
	Pipe: Pipe{
		Handler: formatJSON,
	},
	Output: Replace,
	Tags:   []string{"json", "format", "prettify", "pretty"},
}
