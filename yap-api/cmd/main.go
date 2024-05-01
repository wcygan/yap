package main

import (
	auth_pb "github.com/wcygan/yap/generated/go/auth/v1"
	chat_pb "github.com/wcygan/yap/generated/go/chat/v1"
	"github.com/wcygan/yap/yap-api/internal/auth"
	"github.com/wcygan/yap/yap-api/internal/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	// Register the authentication service
	auth_pb.RegisterAuthServiceServer(s, auth.NewAuthService())

	// Register the chat service
	chat_pb.RegisterClientStreamingServiceServer(s, chat.NewChatService())

	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on %s", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
