package auth

import (
	"context"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	auth "github.com/wcygan/yap/generated/go/auth/v1"
	"path/filepath"
	"testing"
	"time"
)

func setupTestDB(t *testing.T) (context.Context, string) {
	ctx := context.Background()

	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithInitScripts(filepath.Join("..", "..", "..", "auth-db", "init-auth-db.sql")),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	return ctx, connStr
}

// TestAuthService_Register tests that users are only able to register with unique usernames
func TestAuthService_Register(t *testing.T) {
	ctx, connStr := setupTestDB(t)
	authSvc, err := NewAuthService(connStr)
	assert.NoError(t, err)

	tests := []struct {
		name    string
		req     *auth.RegisterRequest
		wantErr bool
	}{
		{
			name: "FirstValidUser",
			req:  &auth.RegisterRequest{Username: "test1", Password: "password"},
		},
		{
			name:    "DuplicateUser",
			req:     &auth.RegisterRequest{Username: "test1", Password: "password"},
			wantErr: true,
		},
		{
			name: "SecondValidUser",
			req:  &auth.RegisterRequest{Username: "test2", Password: "password"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := authSvc.Register(ctx, tt.req)
			assert.Equal(t, tt.wantErr, err != nil)
			if !tt.wantErr {
				assert.NotEmpty(t, resp.AccessToken)
				assert.NotEmpty(t, resp.RefreshToken)
			}
		})
	}
}

// TestLoginAndValidate that users are able to log in and receive valid access and refresh tokens
func TestLoginAndValidate(t *testing.T) {
	ctx, connStr := setupTestDB(t)
	authSvc, err := NewAuthService(connStr)
	assert.NoError(t, err)

	// Register a new user
	c, err := authSvc.Register(ctx, &auth.RegisterRequest{
		Username: "Henry",
		Password: "Password123",
	})

	// The registration should be successful
	assert.NoError(t, err)
	assert.NotNil(t, c)

	// Attempt to log in with the same credentials
	loginResponse, err := authSvc.Login(ctx, &auth.LoginRequest{
		Username: "Henry",
		Password: "Password123",
	})

	// The login should be successful
	assert.NoError(t, err)
	assert.NotNil(t, loginResponse)

	// The access token should be a valid UUID
	err = uuid.Validate(loginResponse.AccessToken)
	assert.NoError(t, err)

	// The refresh token should be a valid UUID
	err = uuid.Validate(loginResponse.RefreshToken)
	assert.NoError(t, err)
}
