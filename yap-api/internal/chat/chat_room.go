package chat

import (
	chat "github.com/wcygan/yap/generated/go/chat/v1"
)

type ChatRoomService struct {
	natsClientUrl string
	chat.UnimplementedChatRoomServiceServer
}

func NewChatRoomService(natsClientUrl string) *ChatRoomService {
	return &ChatRoomService{
		natsClientUrl: natsClientUrl,
	}
}

func (s *ChatRoomService) JoinChatRoom(req *chat.JoinChatRequest, stream chat.ChatRoomService_JoinChatRoomServer) error {
	// TODO: Send a "user joined" message to clients

	// TODO start a nats consumer that will send messages to the client

	//err := stream.Send(temp)
	//if err != nil {
	//	log.Printf("Error sending message: %v", err)
	//	return err
	//} else {
	//	log.Printf("Sent user joined message: %v", temp)
	//}

	<-stream.Context().Done()
	return nil
}
