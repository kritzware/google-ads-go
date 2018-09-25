package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	services "github.com/kritzware/google-ads-go/services"
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

func (g *GoogleAdsClient) GetService(service string) reflect.Value {
	err := g.createGrpcConnection()
	if err != nil {
		panic(err)
	}
	var googleAdsServices services.GoogleAdsServices

	serviceName := fmt.Sprintf("New%sClient", service)
	serviceParams := []reflect.Value{reflect.ValueOf(g.conn)}

	relService := reflect.ValueOf(&googleAdsServices).MethodByName(serviceName)
	return relService.Call(serviceParams)[0]
}

// headers := metadata.Pairs(
// 	"Authorization", fmt.Sprintf("Bearer %s", client.Token.AccessToken),
// 	"developer-token", client.DeveloperToken,
// )
// ctx := metadata.NewOutgoingContext(context.Background(), headers)

func NewGoogleAdsClient(args *GoogleAdsClientArgs) *GoogleAdsClient {
	credentials := &oauth2.Config{
		ClientID:     args.ClientID,
		ClientSecret: args.ClientSecret,
		Endpoint:     google.Endpoint,
	}
	initialToken := &oauth2.Token{
		RefreshToken: args.RefreshToken,
	}
	return &GoogleAdsClient{
		credentials:    credentials,
		token:          initialToken,
		developerToken: args.DeveloperToken,
	}
}

type GoogleAdsStorageConfig struct {
	ClientID       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	RefreshToken   string `json:"refresh_token"`
	DeveloperToken string `json:"developer_token"`
}

func NewGoogleAdsClientFromStorage(filepath string) (*GoogleAdsClient, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var g GoogleAdsStorageConfig
	json.Unmarshal(file, &g)

	client := NewGoogleAdsClient(&GoogleAdsClientArgs{
		ClientID:       g.ClientID,
		ClientSecret:   g.ClientSecret,
		RefreshToken:   g.RefreshToken,
		DeveloperToken: g.DeveloperToken,
	})
	return client, nil
}
