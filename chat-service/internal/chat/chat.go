package chat

import (
	"context"
	chat_pb "github.com/wcygan/yap/generated/go/chat/v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type ChatService struct {
	chat_pb.UnimplementedMessagePersistenceServiceServer
}

func NewChatService() *ChatService {
	return &ChatService{}
}

func (s *ChatService) PersistChatMessage(ctx context.Context, req *chat_pb.MessageToPersist) (*emptypb.Empty, error) {
	log.Printf("Persisting message: %s", req.String())

	// TODO 1: Spin up ScyllaDB instances in Kubernetes
	// TODO 2: Connect to ScyllaDB & persist messages
	// TODO 3: Spin up Pulsar instances in Kubernetes
	// TODO 4: Connect to Pulsar & publish messages

	return &emptypb.Empty{}, nil
}
