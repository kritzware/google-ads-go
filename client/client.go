package client

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address        string = "googleads.googleapis.com:443"
	defaultVersion string = "v0"
)

type GoogleAdsClient struct {
	credentials    *oauth2.Config
	token          *oauth2.Token
	conn           *grpc.ClientConn
	developerToken string
}

type GoogleAdsClientArgs struct {
	ClientID       string
	ClientSecret   string
	DeveloperToken string
	RefreshToken   string
}

type googleAdsStorageConfig struct {
	ClientID       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	RefreshToken   string `json:"refresh_token"`
	DeveloperToken string `json:"developer_token"`
}

func (g *GoogleAdsClient) refreshToken() error {
	tok := g.credentials.TokenSource(context.Background(), g.token)
	newToken, err := tok.Token()
	if err != nil {
		return err
	}
	g.token = newToken
	return nil
}

func (g *GoogleAdsClient) createGrpcConnection() error {
	// TODO: Also pass token creds
	// creds := oauth.NewOauthAccess(client.Token)
	// grpc.WithPerRPCCredentials(creds)
	err := g.refreshToken()
	if err != nil {
		return err
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		return err
	}
	g.conn = conn
	return nil
}

// TODO: Just change this to a public field
// Conn returns a pointer to the clients gRPC connection
func (g *GoogleAdsClient) Conn() *grpc.ClientConn {
	return g.conn
}

// TODO: Add metadata
// headers := metadata.Pairs(	// NewGoogleAdsClient creates a new client with specified credentials
// 	"Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken),
// 	"developer-token", client.DeveloperToken,
// )
// ctx := metadata.NewOutgoingContext(context.Background(), headers)

// NewGoogleAdsClient creates a new client with specified credentials
func NewGoogleAdsClient(args *GoogleAdsClientArgs) *GoogleAdsClient {
	credentials := &oauth2.Config{
		ClientID:     args.ClientID,
		ClientSecret: args.ClientSecret,
		Endpoint:     google.Endpoint,
	}
	initialToken := &oauth2.Token{
		RefreshToken: args.RefreshToken,
	}
	c := &GoogleAdsClient{
		credentials:    credentials,
		token:          initialToken,
		developerToken: args.DeveloperToken,
	}
	c.createGrpcConnection()
	return c
}

// NewGoogleAdsClientFromStorage creates a new client instance from specified "google-ads.json" file
func NewGoogleAdsClientFromStorage(filepath string) (*GoogleAdsClient, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var g googleAdsStorageConfig
	json.Unmarshal(file, &g)

	client := NewGoogleAdsClient(&GoogleAdsClientArgs{
		ClientID:       g.ClientID,
		ClientSecret:   g.ClientSecret,
		RefreshToken:   g.RefreshToken,
		DeveloperToken: g.DeveloperToken,
	})
	return client, nil
}
