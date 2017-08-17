package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/fizzwu/bmwg/chap_01/rpc_http_json/contract"
)

func PerformRequest() contract.HelloWorldResponse {
	r, err := http.Post(
		"Http://localhost:4000",
		"application/json",
		bytes.NewBuffer([]byte(`{"id":1, "method":"HelloWorldHandler.HelloWorld", "params":[{"name":"Sam"}]}`)),
	)
	if err != nil {
		log.Fatal("err:", err)
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	var response contract.HelloWorldResponse
	decoder.Decode(&response)
	return response
}
