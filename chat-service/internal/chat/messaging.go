package chat

import (
	"context"
	chat "github.com/wcygan/yap/generated/go/chat/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MessagingService struct {
	// TODO: add Pulsar producer connection
	// TODO: add ScyllaDB connection
	chat.UnimplementedMessagingServiceServer
}

func NewMessagingService() *MessagingService {
	return &MessagingService{}
}

func (s *MessagingService) SendMessage(ctx context.Context, req *chat.ChatMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
