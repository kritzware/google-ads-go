package auth

import (
	"os"
	"testing"
	"time"

	"google.golang.org/grpc/connectivity"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	clientID       = os.Getenv("GADS_CLIENT_ID")
	clientSecret   = os.Getenv("GADS_CLIENT_SECRET")
	developerToken = os.Getenv("GADS_DEVELOPER_TOKEN")
	refreshToken   = os.Getenv("GADS_REFRESH_TOKEN")
)

func TestNewCredentials(t *testing.T) {
	creds := NewCredentials(clientID, clientSecret)
	if creds.ClientID != clientID {
		t.Fatalf("Credentials client id does not match")
	}
	if creds.ClientSecret != clientSecret {
		t.Fatalf("Credentials client secret does not match")
	}
}

func TestNewPartialToken(t *testing.T) {
	token := NewPartialToken(refreshToken)
	if token.AccessToken != "" {
		t.Fatal("Partial token should have no access token")
	}
	if token.RefreshToken != refreshToken {
		t.Fatal("Refresh token does not match partial token")
	}
}

func TestRefreshToken(t *testing.T) {
	creds := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
	}
	partialToken := NewPartialToken(refreshToken)

	token, _ := RefreshToken(creds, partialToken)
	if !token.Valid() {
		t.Fatalf("Token should be valid after refresh. expected=true, got=%t", token.Valid())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(50 * time.Millisecond):
		break
	case <-ctx.Done():
		break
	}

	token, _ = RefreshToken(creds, partialToken, ctx)
	if token != nil {
		t.Fatal("Using an optional context with timeout should cause token refresh to halt")
	}
}

func TestNewGrpcConnection(t *testing.T) {
	token, _ := createTestToken()
	conn, _, err := NewGrpcConnection(token, developerToken)
	if err != nil {
		t.Fatalf("gRPC connection should not error when using valid credentials. got=%s", err.Error())
	}
	if conn.GetState() != connectivity.Idle {
		t.Fatalf("gRPC connection should be IDLE upon creation. expected=IDLE, got=%s", conn.GetState())
	}
}

func TestCreateHeaders(t *testing.T) {
	token, _ := createTestToken()
	headers := createHeaders(token.AccessToken, developerToken)

	if headers.Len() != 2 {
		t.Fatalf("Metadata headers should have 2 items. expected=2, got=%d", headers.Len())
	}
	if headers.Get("Authorization")[0] == "" {
		t.Fatalf("Metadata headers must have an authorization key")
	}
	if headers.Get("developer-token")[0] != developerToken {
		t.Fatalf("Metadata header \"developer-token\" does not match specified developerToken")
	}
}

func createTestToken() (*oauth2.Token, error) {
	creds := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
	}
	partialToken := NewPartialToken(refreshToken)
	return RefreshToken(creds, partialToken)
}
