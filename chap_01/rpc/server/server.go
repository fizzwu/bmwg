package server

import "github.com/fizzwu/bmwg/chap_01/rpc/contract"
import "net/rpc"
import "net"
import "fmt"
import "log"

const port = 4000

func main() {
	log.Printf("Server starting on port %v\n", port)
	StartServer()
}

type HelloWorldHandler struct {
}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello" + args.Name
	return nil
}

func StartServer() {
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port %v: %s", port, err))
	}
	defer l.Close()

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}
