package parse

import (
	"encoding/json"
	"net/http"
)

func AsJson[T interface{}](request *http.Request) (err error, payload T) {
	err = json.NewDecoder(request.Body).Decode(&payload)
	return
}