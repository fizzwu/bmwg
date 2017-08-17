package main

import (
	"github.com/fizzwu/bmwg/chap_01/rpc_http_json/server"
)

func main() {
	// To execute a request with this server run the below command on your command line
	// curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}' http://localhost:4000
	server.StartServer()
}
