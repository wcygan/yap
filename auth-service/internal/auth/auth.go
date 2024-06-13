package auth

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	auth "github.com/wcygan/yap/generated/go/auth/v1"
	"log"
	"time"
)

var errTokenExpired = errors.New("token has expired")
var errInvalidCredentials = errors.New("invalid credentials")

type AuthService struct {
	auth.UnimplementedAuthServiceServer
	db *sql.DB
}

func NewAuthService(connStr string) (*AuthService, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &AuthService{db: db}, nil
}

func (s *AuthService) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	log.Printf("Registering user: %v", req.Username)

	// Insert user into the users table
	_, err := s.db.ExecContext(ctx, "INSERT INTO users(username, password) VALUES($1, $2)", req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// Generate access and refresh tokens
	accessToken := uuid.New().String()
	refreshToken := uuid.New().String()

	// Store tokens in the tokens table
	_, err = s.db.ExecContext(ctx, "INSERT INTO tokens(user_id, access_token, refresh_token, expires_at) VALUES((SELECT id FROM users WHERE username = $1), $2, $3, $4)", req.Username, accessToken, refreshToken, time.Now().Add(time.Hour*24))
	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		// TODO: Return the user ID
		UserId: "123",
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	log.Printf("Logging in user: %v", req.Username)

	// Retrieve user from the users table
	var userID int64
	var password string
	err := s.db.QueryRowContext(ctx, "SELECT id, password FROM users WHERE username = $1", req.Username).Scan(&userID, &password)
	if err != nil {
		return nil, err
	}

	// Check password
	if password != req.Password {
		return nil, errInvalidCredentials
	}

	// Generate access and refresh tokens
	accessToken := uuid.New().String()
	refreshToken := uuid.New().String()

	// Store tokens in the tokens table
	_, err = s.db.ExecContext(ctx, "INSERT INTO tokens(user_id, access_token, refresh_token, expires_at) VALUES($1, $2, $3, $4)", userID, accessToken, refreshToken, time.Now().Add(time.Hour*24))
	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		// TODO: Return the user ID
		UserId: "123",
	}, nil
}

func (s *AuthService) Validate(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	// Retrieve token from the tokens table
	var userID int64
	var expiresAt time.Time
	err := s.db.QueryRowContext(ctx, "SELECT user_id, expires_at FROM tokens WHERE access_token = $1", req.AccessToken).Scan(&userID, &expiresAt)
	if err != nil {
		return nil, err
	}

	// Check if token is expired
	if expiresAt.Before(time.Now()) {
		return nil, errTokenExpired
	}

	// Retrieve username from the users table
	var username string
	err = s.db.QueryRowContext(ctx, "SELECT username FROM users WHERE id = $1", userID).Scan(&username)
	if err != nil {
		return nil, err
	}

	return &auth.ValidateResponse{
		Username: username,
	}, nil
}

func (s *AuthService) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	// Retrieve token from the tokens table
	var userID int64
	var expiresAt time.Time
	err := s.db.QueryRowContext(ctx, "SELECT user_id, expires_at FROM tokens WHERE refresh_token = $1", req.RefreshToken).Scan(&userID, &expiresAt)
	if err != nil {
		return nil, err
	}

	// Check if refresh token is expired
	if expiresAt.Before(time.Now()) {
		return nil, errTokenExpired
	}

	// Generate new access token
	accessToken := uuid.New().String()

	// Update access token in the tokens table
	_, err = s.db.ExecContext(ctx, "UPDATE tokens SET access_token = $1, expires_at = $2 WHERE refresh_token = $3", accessToken, time.Now().Add(time.Hour*24), req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &auth.RefreshResponse{
		AccessToken: accessToken,
	}, nil
}
