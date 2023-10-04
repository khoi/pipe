package manifest

import (
	"context"
	"sort"
	"strings"
)

func sortLines(_ context.Context, input *string) (string, error) {
	lines := strings.Split(*input, "\n")
	sort.Strings(lines)
	return strings.Join(lines, "\n"), nil
}

var SortLines = Manifest{
	ID:          "sort-lines",
	Name:        "Sort Lines",
	Description: "Sort lines alphabetically",
	Pipe: Pipe{
		Handler: sortLines,
		Args: []Argument{
			LinesArgument,
		},
	},
	Output: Replace,
	Tags:   []string{"sort", "line", "sort-line", "alphabetically"},
}
