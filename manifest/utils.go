package manifest

import (
	"encoding/json"
	"log"
)

func PrettifyJSON(input *string) (string, error) {
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
