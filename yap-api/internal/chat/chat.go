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
	// TODO 1: Figure out how to handle send & recv in multiplex fashion (maybe use select?)
	// TODO 2: Use a channel and add it to an internal connection registry
	//         See https://stackoverflow.com/a/49877632 or https://eli.thegreenplace.net/2020/pubsub-using-channels-in-go/
	// TODO 3: When the connection is dropped, remove it from the registry
	// TODO 4: Create a Pulsar listener in the background that listens for messages and sends them over the right channels

	// TODO: Handle client connection, use Pulsar to fan out messages
	for {
		in, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("Received message: %s", in.String())

		hello := &chat_pb.ChatPacket{PacketType: &chat_pb.ChatPacket_ChatMessage{
			ChatMessage: &chat_pb.ChatMessage{
				Message: "Hello, world!",
			},
		}}

		err = stream.Send(hello)
		if err != nil {
			return err
		}
		log.Printf("Sent message: %s", hello.String())
	}
}
