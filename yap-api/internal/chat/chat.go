package chat

import (
	chat_pb "github.com/wcygan/yap/generated/go/chat/v1"
	"log"
)

type ChatService struct {
	chat_pb.UnimplementedClientStreamingServiceServer
}

func NewChatService() *ChatService {
	return &ChatService{}
}

func (s *ChatService) ChatStream(stream chat_pb.ClientStreamingService_ChatStreamServer) error {
	// TODO: Handle client connection, use Pulsar to fan out messages
	for {
		in, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("Received message: %s", in.String())
	}
}
