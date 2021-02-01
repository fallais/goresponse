package goresponse

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

// JSON writes a JSON response into the ResponseWriter with status code.
func JSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	// Marshal the data
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Set content type header
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Write status
	w.WriteHeader(statusCode)

	// Write response
	w.Write(body)

	return nil
}
