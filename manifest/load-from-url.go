package manifest

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func loadFromURL(_ context.Context, input *string) (string, error) {
	// trim input
	*input = strings.TrimSpace(*input)
	if input == nil || len(*input) == 0 {
		return "", errors.New("input url string is nil")
	}
	resp, err := http.Get(*input)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return "", err
	}

	return string(body), nil
}

var LoadFromURL = Manifest{
	ID:          "load-from-url",
	Name:        "Load From URL",
	Description: "Load the content of a URL",
	Pipe: Pipe{
		Handler: loadFromURL,
		Args: []Argument{
			ContentArgument,
		},
	},
	Output: Replace,
	Tags:   []string{"load", "url", "load-url", "remote", "fetch", "http"},
}
