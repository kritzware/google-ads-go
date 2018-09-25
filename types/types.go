package types

import (
	services "github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/services"
	"google.golang.org/grpc"
)

// GoogleAdsServiceClient

// var GoogleAdsService = map[string]func(*grpc.ClientConn) services.GoogleAdsServiceClient{
// 	"NewGoogleAdsServiceClient": services.NewGoogleAdsServiceClient,
// }

type ClientServices struct {
}

func (cs *ClientServices) NewGoogleAdsServiceClient(conn *grpc.ClientConn) services.GoogleAdsServiceClient {
	return services.NewGoogleAdsServiceClient(conn)
}
