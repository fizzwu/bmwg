package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/fizzwu/bmwg/chap_01/rpc/contract"
)

const port = 4000

type HelloWorldHandler struct {
}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello" + args.Name
	return nil
}

func StartServer() {
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)
	rpc.HandleHTTP() // key different from ServeConn method

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port %v: %s", port, err))
	}
	log.Printf("Server starting on port: %v\n", port)
	http.Serve(l, nil)
}
