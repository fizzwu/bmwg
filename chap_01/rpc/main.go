package main

import (
	"fmt"

	"github.com/fizzwu/bmwg/chap_01/rpc/client"
	"github.com/fizzwu/bmwg/chap_01/rpc/server"
)

func main() {
	go server.StartServer()
	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
