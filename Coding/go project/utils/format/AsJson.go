package format

import "encoding/json"

type Response struct {
	Data interface{} `json:"data"`
	Err  error       `json:"error"`
}

func AsJson(data interface{}, err error) string {
	bytes, _ := json.MarshalIndent(Response{
		Data: data,
		Err:  err,
	}, "", "    ")
	return string(bytes)
}
