package auth

import (
	"context"

	auth "github.com/wcygan/yap/generated/go/auth/v1"
)

type AuthService struct {
	auth.UnimplementedAuthServiceServer
	// Add any necessary fields
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	return &auth.RegisterResponse{
		AccessToken:  "dummy_access_token",
		RefreshToken: "dummy_refresh_token",
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return &auth.LoginResponse{
		AccessToken:  "dummy_access_token",
		RefreshToken: "dummy_refresh_token",
	}, nil
}

func (s *AuthService) Validate(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	return &auth.ValidateResponse{
		Username: "dummy_username",
	}, nil
}

func (s *AuthService) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	return &auth.RefreshResponse{
		AccessToken: "dummy_access_token",
	}, nil
}
