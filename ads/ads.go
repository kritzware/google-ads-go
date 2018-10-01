package ads

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
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
	ctx            context.Context
}

type GoogleAdsClientParams struct {
	ClientID       string
	ClientSecret   string
	DeveloperToken string
	RefreshToken   string
}

type googleAdsStorageParams struct {
	ClientID       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	RefreshToken   string `json:"refresh_token"`
	DeveloperToken string `json:"developer_token"`
}

// NewClient creates a new client with specified credentials
func NewClient(args *GoogleAdsClientParams) (*GoogleAdsClient, error) {
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
	err := c.createGrpcConnection()
	if err != nil {
		return nil, err
	}
	return c, nil
}

// NewClientFromStorage creates a new client instance from specified "google-ads.json" file
func NewClientFromStorage(filepath string) (*GoogleAdsClient, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var g googleAdsStorageParams
	json.Unmarshal(file, &g)

	client, err := NewClient(&GoogleAdsClientParams{
		ClientID:       g.ClientID,
		ClientSecret:   g.ClientSecret,
		RefreshToken:   g.RefreshToken,
		DeveloperToken: g.DeveloperToken,
	})
	return client, err
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
	// creds := oauth.NewOauthAccess(g.token)
	// grpc.WithPerRPCCredentials(creds)
	err := g.refreshToken()
	if err != nil {
		return err
	}

	headers := metadata.Pairs(
		"Authorization", fmt.Sprintf("Bearer %s", g.token.AccessToken),
		"developer-token", g.developerToken,
	)
	ctx := metadata.NewOutgoingContext(context.Background(), headers)
	g.ctx = ctx

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		return err
	}
	g.conn = conn
	return nil
}

// Conn returns a pointer to the clients gRPC connection
func (g *GoogleAdsClient) Conn() *grpc.ClientConn {
	return g.conn
}

// Context returns the context of the client
func (g *GoogleAdsClient) Context() context.Context {
	return g.ctx
}
