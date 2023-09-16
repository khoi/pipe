package manifest

import (
	b64 "encoding/base64"

	"golang.org/x/net/context"
)

func base64Decode(_ context.Context, input *string) (string, error) {
	decoded, err := b64.StdEncoding.DecodeString(*input)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

var Base64Decode = Manifest{
	ID:          "base64-decode",
	Name:        "Base64 Decode",
	Description: "Base64 decode document",
	Pipe: Pipe{
		Handler: base64Decode,
	},
	Output: Replace,
	Tags:   []string{"base64", "decode", "base64-decode"},
}
