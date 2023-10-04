package manifest

import (
	"golang.org/x/net/context"
)

func formatJSON(_ context.Context, input *string) (string, error) {
	return PrettifyJSON(input)
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
