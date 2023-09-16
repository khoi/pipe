package manifest

import (
	b64 "encoding/base64"

	"golang.org/x/net/context"
)

func base64Encode(_ context.Context, input *string) (string, error) {
	encoded := b64.StdEncoding.EncodeToString([]byte(*input))
	return encoded, nil
}

var Base64Encode = Manifest{
	ID:          "base64-encode",
	Name:        "Base64 Encode",
	Description: "Base64 encode document",
	Pipe: Pipe{
		Handler: base64Encode,
	},
	Output: Replace,
	Tags:   []string{"base64", "encode", "base64-encode"},
}
