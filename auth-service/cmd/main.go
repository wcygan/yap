package main

import (
	"github.com/wcygan/yap/auth-service/internal/auth"
	authpb "github.com/wcygan/yap/generated/go/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	svc, err := auth.NewAuthService("postgres://postgres:your-password-here@auth-db:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to create auth service: %v", err)
	}

	s := grpc.NewServer()

	reflection.Register(s)
	log.Printf("reflection is enabled")

	authpb.RegisterAuthServiceServer(s, svc)
	log.Printf("auth service is registered")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("auth-service is listening on %s", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
