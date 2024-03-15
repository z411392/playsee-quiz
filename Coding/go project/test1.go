package main

import (
	"fmt"
	"net/http"
)

func Test1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test 1:")
}
