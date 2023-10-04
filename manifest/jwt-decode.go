package manifest

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/context"
)

func jwtDecode(_ context.Context, input *string) (string, error) {
	s := strings.Split(*input, ".")

	if len(s) != 3 {
		return "", errors.New("invalid JWT structure")
	}

	header, err := base64.RawURLEncoding.DecodeString(s[0])
	if err != nil {
		return "", errors.New("decode header failed")
	}

	payload, err := base64.RawURLEncoding.DecodeString(s[1])
	if err != nil {
		return "", errors.New("decode payload failed")
	}

	signature := s[2]

	decodedStr := fmt.Sprintf(
		`{"header":%s,"payload":%s,"signature":"%s"}`,
		header,
		payload,
		signature,
	)

	return PrettifyJSON(&decodedStr)
}

var JwtDecode = Manifest{
	ID:          "jwt-decode",
	Name:        "JWT Decode",
	Description: "Decode a JWT token",
	Pipe: Pipe{
		Handler: jwtDecode,
	},
	Output: Replace,
	Tags:   []string{"jwt", "decode"},
}
