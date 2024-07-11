package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"sync"
	"test_protobuf/chat"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	api := flag.String("api", "sayHello", "Select a server to run the API")
	flag.Parse()

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	if *api == "sayHello" {
		messages := []string{"Jirawan", "Happy", "Charlie", "Thomas"}
		var wg sync.WaitGroup
		for _, msg := range messages {
			wg.Add(1)
			go makeRequest(c, msg, &wg)
		}
		wg.Wait()
	} else if *api == "streamData" {
		stream, err := c.StreamData(context.Background(), &chat.StreamDataRequest{Message: "Start pushing"})
		if err != nil {
			log.Fatalf("could not push: %v", err)
		}

		for {
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					fmt.Println("Server finished sending messages.")
					break
				}
				log.Fatalf("error receiving: %v", err)
			}
			fmt.Printf("Received: %s\n", res.Data)
		}
	}

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
