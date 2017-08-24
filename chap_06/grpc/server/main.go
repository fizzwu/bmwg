package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	proto "github.com/fizzwu/bmwg/chap_06/grpc/proto"
	context "golang.org/x/net/context"
)

type kittenServer struct {
}

func (ks *kittenServer) Hello(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	response := &proto.Response{}
	response.Msg = fmt.Sprintf("Hello, %v", request.Name)

	return response, nil
}

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterKittensServer(grpcServer, &kittenServer{})
	grpcServer.Serve(ln)
}
