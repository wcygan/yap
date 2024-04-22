package auth

import (
	"context"
	auth "github.com/wcygan/yap/generated/go/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

type AuthService struct {
	auth.UnimplementedAuthServiceServer
}

var authClient auth.AuthServiceClient

func init() {
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
	authClient = auth.NewAuthServiceClient(conn)
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	return authClient.Register(ctx, req)
}

func (s *AuthService) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return authClient.Login(ctx, req)
}

func (s *AuthService) Validate(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	return authClient.Validate(ctx, req)
}

func (s *AuthService) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	return authClient.Refresh(ctx, req)
}
