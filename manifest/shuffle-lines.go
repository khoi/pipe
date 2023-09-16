package manifest

import (
	"context"
	"math/rand"
	"strings"
)

func shuffleLines(_ context.Context, input *string) (string, error) {
	lines := strings.Split(*input, "\n")
	rand.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})

	return strings.Join(lines, "\n"), nil
}

var ShuffleLines = Manifest{
	ID:          "shuffle-lines",
	Name:        "Shuffle Lines",
	Description: "Shuffle lines of the document",
	Pipe: Pipe{
		Handler: shuffleLines,
		Args: []Argument{
			LinesArgument,
		},
	},
	Output: Replace,
	Tags:   []string{"shuffle", "line", "shuffle-line", "randomize"},
}
