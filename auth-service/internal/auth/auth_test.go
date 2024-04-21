package auth

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	auth "github.com/wcygan/yap/generated/go/auth/v1"
	"testing"
)

func setupTestDatabase(t *testing.T) *sql.DB {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15.3",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %v", err)
		}
	}()

	host, err := container.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}

	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatal(err)
	}

	connStr := fmt.Sprintf("postgres://postgres:password@%s:%s/testdb?sslmode=disable", host, port.Port())

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestRegister(t *testing.T) {
	db := setupTestDatabase(t)
	defer db.Close()

	s := &AuthService{db: db}

	// Create the necessary tables for testing
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username VARCHAR(255) UNIQUE NOT NULL,
            password VARCHAR(255) NOT NULL
        );

        CREATE TABLE IF NOT EXISTS tokens (
            user_id INTEGER REFERENCES users(id),
            access_token VARCHAR(255) UNIQUE NOT NULL,
            refresh_token VARCHAR(255) UNIQUE NOT NULL,
            expires_at TIMESTAMP NOT NULL
        );
    `)
	if err != nil {
		t.Fatal(err)
	}

	// Test the Register method
	_, err = s.Register(context.Background(), &auth.RegisterRequest{
		Username: "testuser",
		Password: "testpassword",
	})

	if err != nil {
		t.Fatal(err)
	}
}
