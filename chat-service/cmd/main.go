package main

import (
	"github.com/wcygan/yap/chat-service/internal/chat"
	"github.com/wcygan/yap/chat-service/internal/storage"
	chatpb "github.com/wcygan/yap/generated/go/chat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	chatDB, err := storage.NewStorage("chat-db")
	if err != nil {
		log.Fatalf("failed to create storage: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	log.Printf("reflection is enabled")

	chatService, err := chat.NewChatService(chatDB)
	if err != nil {
		log.Fatalf("failed to create chat service: %v", err)
	}

	// Register the chat service
	chatpb.RegisterMessagePersistenceServiceServer(s, chatService)
	log.Printf("message persistence service is registered")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("chat-service is listening on %s", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
