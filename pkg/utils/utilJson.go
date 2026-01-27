package utils

import (
	"encoding/json"
)

func WriteJson[T any](data T) (string, []byte, error) {
	var contentType string = "application/json"

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "Error encoding to JSON", nil, err
	}

	return contentType, jsonData, nil
}
