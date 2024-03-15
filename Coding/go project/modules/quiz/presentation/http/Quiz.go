package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	queries "playsee.co/interview/modules/quiz/application/queries"
)

func Quiz(responseWriter http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body) // https://stackoverflow.com/questions/15672556/handling-json-post-request-in-go
	payload := &struct {                     // https://blog.boot.dev/golang/anonymous-structs-golang/
		Array []interface{}
	}{}
	decoder.Decode(payload)
	linkedList := queries.ReadAsLinkedList(payload.Array...)
	fmt.Fprintf(responseWriter, "Test 1: %v", linkedList)
}
