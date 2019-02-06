package ads

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var (
	clientID       = os.Getenv("GADS_CLIENT_ID")
	clientSecret   = os.Getenv("GADS_CLIENT_SECRET")
	developerToken = os.Getenv("GADS_DEVELOPER_TOKEN")
	refreshToken   = os.Getenv("GADS_REFRESH_TOKEN")
)

var params = GoogleAdsClientParams{
	ClientID:       clientID,
	ClientSecret:   clientSecret,
	DeveloperToken: developerToken,
	RefreshToken:   refreshToken,
}

func TestNewClient(t *testing.T) {
	if _, err := NewClient(&GoogleAdsClientParams{
		ClientID:       "",
		ClientSecret:   "",
		DeveloperToken: "",
		RefreshToken:   "",
	}); err == nil {
		t.Fatal("NewClient should return an error when GoogleAdsClientParams are nil")
	}
	if _, err := NewClient(&GoogleAdsClientParams{}); err == nil {
		t.Fatal("NewClient should return an error when GoogleAdsClientParams are undefined")
	}

	client, err := NewClient(&params)
	if err != nil {
		fmt.Println(err.Error())
		t.Fatal("NewClient should not return an error when using valid credentials")
	}

	conn := client.Conn()
	address := "googleads.googleapis.com:443"

	if conn == nil {
		t.Fatalf("client.Conn should return a new grpc.ClientConn. got=%+v", conn)
	}
	if target := conn.Target(); target != address {
		t.Fatalf("NewClient connection address incorrect. expected=%s, got=%s", address, target)
	}
	if validToken := client.TokenIsValid(); validToken == false {
		t.Fatalf("NewClient created an invalid token with valid credentials. expected=true, got=%t", validToken)
	}
}

func TestNewClientErrors(t *testing.T) {
	_, err := NewClient(&GoogleAdsClientParams{
		ClientID:       "abc",
		ClientSecret:   "123",
		DeveloperToken: "abc",
		RefreshToken:   "123",
	})
	if !strings.Contains(err.Error(), "401 Unauthorized") {
		t.Fatalf("NewClient should return 401 Unauthorized when no Client credentials are specified. expected=%s, got=%s", "401 Unauthorized", err.Error())
	}
}

func TestUsingLoginCustomerID(t *testing.T) {
	client, _ := NewClient(&GoogleAdsClientParams{
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		DeveloperToken:  developerToken,
		RefreshToken:    refreshToken,
		LoginCustomerID: "123",
	})
	if client.loginCustomerID != "123" {
		t.Fatalf("client.loginCustomerID should be used when specified.")
	}
}

func TestNewClientFromStorage(t *testing.T) {
	_, err := NewClientFromStorage("./auth_test.json")
	if err == nil {
		t.Fatal("NewClientFromStorage should return an error when using invalid oauth credentials")
	}
}
