package main

import (
	"log"
	"sync"
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

	messages := []string{"Jirawan", "Happy", "Charlie", "Thomas"}
	var wg sync.WaitGroup
	for _, msg := range messages {
		wg.Add(1)
		go makeRequest(c, msg, &wg)
	}
	wg.Wait()

}

func makeRequest(c chat.ChatServiceClient, msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	message := chat.Message{
		Body: msg,
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("response: ", response)
}
