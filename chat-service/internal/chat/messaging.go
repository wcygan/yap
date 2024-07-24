package chat

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/wcygan/yap/chat-service/internal/persistence"
	chat "github.com/wcygan/yap/generated/go/chat/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type MessagingService struct {
	// TODO: add Pulsar producer connection
	storage *persistence.Storage
	chat.UnimplementedMessagingServiceServer
}

func NewMessagingService(storage *persistence.Storage) *MessagingService {
	return &MessagingService{
		storage: storage,
	}
}

func (s *MessagingService) SendMessage(ctx context.Context, req *chat.ChatMessage) (*emptypb.Empty, error) {
	// Convert the request's data to appropriate types
	timestamp := time.UnixMilli(req.Timestamp)
	userId, err := gocql.ParseUUID(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID: %v", err)
	}

	// Persist the message
	err = s.storage.SaveMessage(req.ChannelName, userId, req.Message, timestamp)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to save message: %v", err)
	}

	// TODO: produce the message to Pulsar
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
