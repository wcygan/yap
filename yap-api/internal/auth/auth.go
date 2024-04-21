package auth

import (
	"context"
	auth "github.com/wcygan/yap/generated/go/auth/v1"
)

type AuthService struct {
	auth.UnimplementedAuthServiceServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	return &auth.RegisterResponse{
		AccessToken:  "test",
		RefreshToken: "test",
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return &auth.LoginResponse{
		AccessToken:  "test",
		RefreshToken: "test",
	}, nil
}

func (s *AuthService) Validate(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	return &auth.ValidateResponse{
		Username: "test",
	}, nil
}

func (s *AuthService) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	return &auth.RefreshResponse{
		AccessToken: "test",
	}, nil
}
