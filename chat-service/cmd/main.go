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

	// Register the chat service
	chat_pb.RegisterMessagePersistenceServiceServer(s, chat.NewChatService())

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on %s", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
