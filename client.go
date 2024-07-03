package main

import (
	"log"
	"test_protobuf/chat"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	message := chat.Message{
		Body: "Jirwan",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("response: ", response)
}
