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

// TestAuthenticationLifecycle asserts that users can login with their credentials & validate + refresh their access token
func TestAuthenticationLifecycle(t *testing.T) {
	ctx, connStr := setupTestDB(t)
	authSvc, err := NewAuthService(connStr)
	assert.NoError(t, err)

	// (1) Register a new user
	c, err := authSvc.Register(ctx, &auth.RegisterRequest{
		Username: "Henry",
		Password: "Password123",
	})

	// The registration should be successful
	assert.NoError(t, err)
	assert.NotNil(t, c)

	// (2) Attempt to log in with the same credentials
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

	// (3) It should be possible to validate the access token
	validateResponse1, err := authSvc.Validate(ctx, &auth.ValidateRequest{
		AccessToken: loginResponse.AccessToken,
	})

	// The validation should be successful
	assert.NoError(t, err)

	// The user ID should be the same as the one used to log in
	assert.Equal(t, "Henry", validateResponse1.Username)

	// (4) Test the refresh token
	refreshResponse, err := authSvc.Refresh(ctx, &auth.RefreshRequest{
		RefreshToken: loginResponse.RefreshToken,
	})

	// The refresh should be successful
	assert.NoError(t, err)

	// The access token should be a valid UUID
	err = uuid.Validate(refreshResponse.AccessToken)
	assert.NoError(t, err)

	// The new access token (from the refresh) should be different from the old access token
	assert.NotEqual(t, loginResponse.AccessToken, refreshResponse.AccessToken)

	// (5) It should be possible to validate the new access token
	validateResponse2, err := authSvc.Validate(ctx, &auth.ValidateRequest{
		AccessToken: refreshResponse.AccessToken,
	})

	// The validation should be successful
	assert.NoError(t, err)

	// The user ID should be the same as the one used to log in
	assert.Equal(t, "Henry", validateResponse2.Username)
}
