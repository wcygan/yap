package main

import (
	"log"
	"net"

	"github.com/wcygan/yap/auth-service/internal/auth"
	auth_pb "github.com/wcygan/yap/generated/go/auth/v1"
	"google.golang.org/grpc"
)

func main() {
	svc, err := auth.NewAuthService("postgres://postgres:your-password-here@auth-db:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to create auth service: %v", err)
	}

	s := grpc.NewServer()
	auth_pb.RegisterAuthServiceServer(s, svc)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on %s", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
