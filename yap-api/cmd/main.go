package main

import (
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	authpb "github.com/wcygan/yap/generated/go/auth/v1"
	chatpb "github.com/wcygan/yap/generated/go/chat/v1"
	"github.com/wcygan/yap/yap-api/internal/auth"
	"github.com/wcygan/yap/yap-api/internal/chat"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

func main() {
	nats, err := createNatsServer()
	if err != nil {
		log.Fatalf("failed to create nats server: %v", err)
	}

	natsClientUrl := nats.ClientURL()

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
	chatpb.RegisterChatRoomServiceServer(s, chat.NewChatRoomService(natsClientUrl))
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

func createNatsServer() (*server.Server, error) {
	opts := &server.Options{}

	// Initialize new natsServer with options
	natsServer, err := server.NewServer(opts)

	if err != nil {
		panic(err)
	}

	// Start the natsServer via goroutine
	go natsServer.Start()

	// Wait for natsServer to be ready for connections
	if !natsServer.ReadyForConnections(4 * time.Second) {
		panic("not ready for connection")
	}

	log.Printf("NATS server is ready")

	return natsServer, nil
}

func startPulsarConsumer(natsClientUrl string) {
	_, err := nats.Connect(natsClientUrl)
	if err != nil {
		log.Fatalf("failed to connect to nats: %v", err)
	}
	// TODO implement pulsar consumer
	// 		Maybe spawn a new goroutine to handle the consumer
}
