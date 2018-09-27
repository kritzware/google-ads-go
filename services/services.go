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

func (s *CampaignService) GetCampaign(cid, id string) (*gaResources.Campaign, error) {
	resourceName := fmt.Sprintf("customers/%s/campaigns/%s", cid, id)
	request := gaServices.GetCampaignRequest{ResourceName: resourceName}
	campaign, err := s.service.GetCampaign(s.ctx, &request)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

type AdGroupService struct {
	service gaServices.AdGroupServiceClient
	ctx     context.Context
}

func NewAdGroupService(conn serviceConn, ctx context.Context) *AdGroupService {
	return &AdGroupService{service: gaServices.NewAdGroupServiceClient(conn), ctx: ctx}
}

func (s *AdGroupService) GetAdGroup(cid, id string) (*gaResources.AdGroup, error) {
	resourceName := fmt.Sprintf("customers/%s/adGroups/%s", cid, id)
	request := gaServices.GetAdGroupRequest{ResourceName: resourceName}
	adGroup, err := s.service.GetAdGroup(s.ctx, &request)
	if err != nil {
		return nil, err
	}
	return adGroup, nil
}
