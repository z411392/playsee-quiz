package http

import (
	"fmt"
	"net/http"

	queries "playsee.co/interview/modules/quiz/application/queries"
	parse "playsee.co/interview/utils/parse"
)

type Payload struct {
	Array []interface{}
}

func Quiz(responseWriter http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	err, payload := parse.AsJson[Payload](request)
	var data string
	if err == nil {
		data, err = queries.ReadAsLinkedList(payload.Array...)
	}
	// response := format.AsJson(data, err)
	// fmt.Fprint(responseWriter, response)
	_ = err
	fmt.Fprintf(responseWriter, "Test 1: %s", data)
}
