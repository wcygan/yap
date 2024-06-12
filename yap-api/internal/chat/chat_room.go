package chat

import (
	chat "github.com/wcygan/yap/generated/go/chat/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChatRoomService struct {
	// TODO: add Pulsar producer connection
	chat.UnimplementedChatRoomServiceServer
}

func NewChatRoomService() *ChatRoomService {
	return &ChatRoomService{}
}

func (s *ChatRoomService) JoinChatRoom(req *chat.JoinChatRequest, stream chat.ChatRoomService_JoinChatRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinChatRoom not implemented")
}
