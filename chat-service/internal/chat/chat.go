package chat

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/wcygan/yap/chat-service/internal/storage"
	chat_pb "github.com/wcygan/yap/generated/go/chat/v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
)

type ChatService struct {
	chat_pb.UnimplementedMessagePersistenceServiceServer
	chatDB *storage.Storage
}

func NewChatService(storage *storage.Storage) (*ChatService, error) {
	return &ChatService{chatDB: storage}, nil
}

func (s *ChatService) PersistChatMessage(ctx context.Context, req *chat_pb.MessageToPersist) (*emptypb.Empty, error) {
	log.Printf("Persisting message: %s", req.String())

	timestamp := time.Unix(req.Timestamp, 0)

	channelId, err := gocql.ParseUUID(req.ChannelId)
	if err != nil {
		log.Printf("Failed to parse channel ID: %v", err)
		return nil, err
	}

	userId, err := gocql.ParseUUID(req.UserId)
	if err != nil {
		log.Printf("Failed to parse user ID: %v", err)
		return nil, err
	}

	err = s.chatDB.SaveMessage(channelId, userId, req.Message, timestamp)
	if err != nil {
		log.Printf("Failed to persist message: %v", err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ChatService) Close() {
	s.chatDB.Close()
}
