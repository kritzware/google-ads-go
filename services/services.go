package services

import (
	"google.golang.org/grpc"

	services "github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/services"
)

type serviceConn *grpc.ClientConn

type GoogleAdsServices struct {
}

func (ga *GoogleAdsServices) NewGoogleAdsServiceClient(conn serviceConn) services.GoogleAdsServiceClient {
	return services.NewGoogleAdsServiceClient(conn)
}
