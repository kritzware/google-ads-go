package client

import (
	"context"

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
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		return err
	}
	g.conn = conn
	return err
}

func (g *GoogleAdsClient) GetService(service string) interface{} {
	err := g.createGrpcConnection()
	if err != nil {
		panic(err)
	}

	// serviceName := fmt.Sprintf("New%sClient", service)

	// var t types.ClientServices
	// adsService := reflect.ValueOf(&t).MethodByName(serviceName).Call([]reflect.Value{reflect.Value(g.conn)})

	// return adsService
	return nil
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
