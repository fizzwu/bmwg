package client

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/fizzwu/bmwg/chap_01/rpc/contract"
)

const port = 4000

func CreateClient() *rpc.Client {
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

func PerformRequest(client *rpc.Client) contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "Sam"}
	var reply contract.HelloWorldResponse

	err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)
	if err != nil {
		log.Fatal("call rpc err:", err)
	}
	return reply
}
