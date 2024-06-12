package chat

import (
	chat "github.com/wcygan/yap/generated/go/chat/v1"
	"log"
	"time"
)

type ChatRoomService struct {
	// TODO: add Pulsar producer connection
	chat.UnimplementedChatRoomServiceServer
}

func NewChatRoomService() *ChatRoomService {
	return &ChatRoomService{}
}

func (s *ChatRoomService) JoinChatRoom(req *chat.JoinChatRequest, stream chat.ChatRoomService_JoinChatRoomServer) error {
	// TODO: produce a Packet_UserJoined message to the Pulsar topic
	temp := &chat.Packet{Contents: &chat.Packet_UserJoined{UserJoined: &chat.UserJoinedMessage{
		ChannelId: req.ChannelName,
		UserId:    req.UserId,
		UserName:  req.UserName,
		Timestamp: time.Now().Unix(),
	}}}

	err := stream.Send(temp)
	if err != nil {
		log.Printf("Error sending message: %v", err)
		return err
	} else {
		log.Printf("Sent user joined message: %v", temp)
	}

	// TODO: add Pulsar consumer connection to receive messages from the chat room topic

	<-stream.Context().Done()
	return nil
}
