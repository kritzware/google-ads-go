package services

import (
	"fmt"

	gaResources "github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/resources"
	gaServices "github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/services"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type serviceConn *grpc.ClientConn

type CampaignService struct {
	service gaServices.CampaignServiceClient
	ctx     context.Context
}

func NewCampaignService(conn serviceConn, ctx context.Context) *CampaignService {
	return &CampaignService{service: gaServices.NewCampaignServiceClient(conn), ctx: ctx}
}

func (s *CampaignService) GetCampaign(cid, id string) gaResources.Campaign {
	resourceName := fmt.Sprintf("customers/%s/campaigns/%s", cid, id)
	request := gaServices.GetCampaignRequest{ResourceName: resourceName}
	campaign, err := s.service.GetCampaign(s.ctx, &request)
	if err != nil {
		panic(err)
	}
	return gaResources.Campaign(campaign)
}
