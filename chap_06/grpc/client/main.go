package main

import (
	"fmt"
	"log"

	proto "github.com/fizzwu/bmwg/chap_06/grpc/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Unable to create connection to server: ", err)
	}

	client := proto.NewKittensClient(conn)
	resp, err := client.Hello(context.Background(), &proto.Request{Name: "Nic"})
	if err != nil {
		log.Fatal("Error calling service:", err)
	}

	fmt.Println(resp.Msg)
}
