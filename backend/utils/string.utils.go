package utils

import (
	"encoding/json"
	"errors"
	"strings"
)

func CSVToJSONArray(input string) (string, error) {
	items := strings.Split(input, ",")
	if len(items) < 1 {
		return "", errors.New("nothing to convert")
	}
	for i := range items {
		items[i] = strings.TrimSpace(items[i])
	}

	jsonBytes, err := json.Marshal(items)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}
