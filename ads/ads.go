package ads

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kritzware/google-ads-go/auth"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
)

const (
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

// NewClient creates a new client with specified credential params
func NewClient(params *GoogleAdsClientParams) (*GoogleAdsClient, error) {
	credentials := auth.NewCredentials(params.ClientID, params.ClientSecret)
	initialToken := auth.NewPartialToken(params.RefreshToken)

	c := &GoogleAdsClient{
		credentials:    credentials,
		token:          initialToken,
		developerToken: params.DeveloperToken,
	}

	newToken, err := auth.RefreshToken(c.credentials, c.token)
	if err != nil {
		return nil, err
	}
	c.token = newToken

	conn, ctx, err := auth.NewGrpcConnection(c.token, c.developerToken)
	if err != nil {
		return nil, err
	}
	c.conn = conn
	c.ctx = ctx

	return c, nil
}

// NewClientFromStorage creates a new client instance from specified "google-ads.json" file
func NewClientFromStorage(filepath string) (*GoogleAdsClient, error) {
	params, err := ReadCredentialsFile(filepath)
	if err != nil {
		return nil, err
	}
	client, err := NewClient(params)
	return client, err
}

// ReadCredentialsFile reads a credentials JSON file and returns the exported config
func ReadCredentialsFile(filepath string) (*GoogleAdsClientParams, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var g googleAdsStorageParams
	json.Unmarshal(file, &g)

	return &GoogleAdsClientParams{
		ClientID:       g.ClientID,
		ClientSecret:   g.ClientSecret,
		RefreshToken:   g.RefreshToken,
		DeveloperToken: g.DeveloperToken,
	}, nil
}

// Conn returns a pointer to the clients gRPC connection
func (g *GoogleAdsClient) Conn() *grpc.ClientConn {
	return g.conn
}

// Context returns the context of the client
func (g *GoogleAdsClient) Context() context.Context {
	return g.ctx
}

// TokenIsValid returns a bool indicating if the generated access token is valid
func (g *GoogleAdsClient) TokenIsValid() bool {
	return g.token.Valid()
}
