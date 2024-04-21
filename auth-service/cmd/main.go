package main

import (
	"log"
	"net"

	"github.com/wcygan/yap/auth-service/internal/auth"
	auth_pb "github.com/wcygan/yap/generated/go/auth/v1"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on %s", lis.Addr().String())
	s := grpc.NewServer()
	auth_pb.RegisterAuthServiceServer(s, auth.NewAuthService())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
