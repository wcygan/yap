package main

import (
	authpb "github.com/wcygan/yap/generated/go/auth/v1"
	chatpb "github.com/wcygan/yap/generated/go/chat/v1"
	"github.com/wcygan/yap/yap-api/internal/auth"
	"github.com/wcygan/yap/yap-api/internal/chat"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	reflection.Register(s)
	log.Printf("reflection is enabled")

	// Register the authentication service
	authpb.RegisterAuthServiceServer(s, auth.NewAuthService())
	log.Printf("auth service is registered")

	// Register the chat service
	chatpb.RegisterMessagingServiceServer(s, chat.NewMessagingService())
	log.Printf("chat service is registered")

	// Register the chat room service
	chatpb.RegisterChatRoomServiceServer(s, chat.NewChatRoomService())
	log.Printf("chat room service is registered")

	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("yap-api is listening on %s", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
