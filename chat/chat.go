package chat

import(
	"log"
	"golang.org/x/net/context"
)

type Server struct {}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
    log.Printf("Received: %v", in.Body)
    return &Message{Body: "Hello " + in.Body}, nil
}