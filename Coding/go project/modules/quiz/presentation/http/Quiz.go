package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	queries "playsee.co/interview/modules/quiz/application/queries"
)

type Payload struct {
	Array []interface{}
}

func Quiz(responseWriter http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	payload := &Payload{}
	err := json.NewDecoder(request.Body).Decode(&payload)
	var data string
	if err == nil {
		data, err = queries.ReadAsLinkedList(payload.Array...)
	}
	// response := format.AsJson(data, err)
	// fmt.Fprint(responseWriter, response)
	_ = err
	fmt.Fprintf(responseWriter, "Test 1: %s", data)
}
