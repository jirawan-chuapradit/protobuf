package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"test_protobuf/chat"
)



func main() {
	listener, err := net.Listen("tcp",":9000")
	if err != nil {
		log.Fatal(err)
	}

	s := chat.Server{}
	grpcServer := grpc.NewServer()
	

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}