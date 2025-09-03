package utils

import (
	"encoding/json"
	"io"
)

func ReadJSON[T any](data io.Reader) (T, error) {
    var entity T
    err := json.NewDecoder(data).Decode(&entity)
    return entity, err
}

func WriteJson[T any](data T) (string, []byte, error) {
	var contentType string = "application/json"

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "Error encoding to JSON", nil, err
	}

	return contentType, jsonData, nil
}
