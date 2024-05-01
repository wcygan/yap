package chat

import (
	"context"
	"log"

	chat_pb "github.com/wcygan/yap/generated/go/chat/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatService struct {
	chat_pb.UnimplementedMessagePersistenceServiceServer
}

func NewChatService() *ChatService {
	return &ChatService{}
}

func (s *ChatService) PersistChatMessage(ctx context.Context, req *chat_pb.MessageToPersist) (*emptypb.Empty, error) {
	// TODO: Persist message to ScyllaDB and publish to Pulsar
	log.Printf("Persisting message: %s", req.String())
	return &emptypb.Empty{}, nil
}
