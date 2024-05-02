package main

import (
	"log"
	"net"

	"github.com/wcygan/yap/chat-service/internal/chat"
	chat_pb "github.com/wcygan/yap/generated/go/chat/v1"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	chatService, err := chat.NewChatService()
	if err != nil {
		log.Fatalf("failed to create chat service: %v", err)
	}

	// Register the chat service
	chat_pb.RegisterMessagePersistenceServiceServer(s, chatService)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on %s", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
