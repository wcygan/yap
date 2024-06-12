package chat

import (
	"context"
	chat "github.com/wcygan/yap/generated/go/chat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
)

type MessagingService struct {
	messaginClient chat.MessagingServiceClient
	chat.UnimplementedMessagingServiceServer
}

func NewMessagingService() *MessagingService {
	conn, err := grpc.Dial("auth-service:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             5 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoff.Config{
				BaseDelay:  1.0 * time.Second,
				Multiplier: 1.5,
				Jitter:     0.2,
				MaxDelay:   60 * time.Second,
			},
			MinConnectTimeout: time.Second * 10,
		}),
	)
	if err != nil {
		log.Fatalf("failed to connect to auth service: %v", err)
	}
	authClient := chat.NewMessagingServiceClient(conn)

	return &MessagingService{
		messaginClient: authClient,
	}
}

func (s *MessagingService) SendMessage(ctx context.Context, req *chat.ChatMessage) (*emptypb.Empty, error) {
	response, err := s.messaginClient.SendMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
