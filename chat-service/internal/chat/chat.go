package chat

import (
	"context"
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

func NewChatService() (*ChatService, error) {
	chatDB, err := storage.NewStorage("chat-db")
	if err != nil {
		return nil, err
	}

	return &ChatService{chatDB: chatDB}, nil
}

func (s *ChatService) PersistChatMessage(ctx context.Context, req *chat_pb.MessageToPersist) (*emptypb.Empty, error) {
	log.Printf("Persisting message: %s", req.String())

	err := s.chatDB.SaveMessage(req.RoomName, req.Message.UserId, req.Message.Message, time.Unix(req.Message.Timestamp, 0))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ChatService) Close() {
	s.chatDB.Close()
}
