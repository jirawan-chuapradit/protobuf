package chat

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Received: %v", in.Body)
	return &Message{Body: "Hello " + in.Body}, nil
}

func (s *Server) StreamData(in *StreamDataRequest, stream ChatService_StreamDataServer) error {
	log.Printf("Received: %v", in.Message)
	// set rule for server push
	for i := 0; i < 10; i++ {
		res := &StreamDataResponse{Data: fmt.Sprintf("Data chunk %d", i)}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(1 * time.Second) // Simulate delay
	}
	stream.Context().Done()
	return nil
}
