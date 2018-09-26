package services

import (
	// gaServices "github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/services"

	"context"
	"fmt"

	// gaServices "google/ads/googleads/v0/services"
	gaServices "github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/services"

	"google.golang.org/grpc"
)

type serviceConn *grpc.ClientConn

type CampaignServiceClient struct {
	service gaServices.CampaignServiceClient
}

func NewCampaignService(conn serviceConn) CampaignServiceClient {
	return CampaignServiceClient{service: gaServices.NewCampaignServiceClient(conn)}
}

// func (c *CampaignServiceClient) GetCampaign(request *gaServices.GetCampaignRequest) int {
func (c *CampaignServiceClient) GetCampaign() int {
	request := gaServices.GetCampaignRequest{ResourceName: "Wasd"}
	// fmt.Printf("%+v\n", request)
	fmt.Println(&request)

	campaign, err := c.service.GetCampaign(context.Background(), &request)
	if err != nil {
		panic(err)
	}
	// // return campaign
	// fmt.Printf("%+v\n", campaign)
	fmt.Println(campaign)

	return 1
}

// var ServiceRegistry = make(map[string]reflect.Type)

// func registerType(typedNil interface{}) {
// 	t := reflect.TypeOf(typedNil).Elem()
// 	ServiceRegistry[t.PkgPath()+"."+t.Name()] = t
// }

// type GoogleAdsServiceClient gaServices.GoogleAdsServiceClient

// func init() {
// 	registerType((GoogleAdsServiceClient)(nil))
// }

// var ServiceRegistry = map[string]func(*grpc.ClientConn, *interface{}){
// 	"CampaignService": NewCampaignServiceClient,
// }

// gaServices "github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/services"

// type ServiceConn *grpc.ClientConn
// var typeRegistry = make(map[string]GoogleAdsService)

// func init() {
// typeRegistry["CampaignService"] = &CampaignService{}
// }

// func makeInstance(name string) interface{} {
// 	v := reflect.New(typeRegistry[name]).Elem()
// 	// Maybe fill in fields here if necessary
// 	return v.Interface()
// }

// func GetService(name string, conn *grpc.ClientConn) interface{} {
// 	// serviceName := fmt.Sprintf("New%sClient", name)
// 	service := typeRegistry[name].New(name, conn)
// 	fmt.Printf("%+v\n", service)

// 	// service := x.New(name)
// 	// fmt.Printf("%+v\n", service)
// 	// service := ga.New(serviceName)
// 	// fmt.Println(service)
// 	return typeRegistry[name]
// }

// type GoogleAdsService interface {
// 	New(name string, conn *grpc.ClientConn) interface{}
// }

// type CampaignService struct {
// 	gaServices.GetCampaignRequest
// 	gaServices.MutateCampaignsRequest
// }

// func (s *CampaignService) New(name string, conn *grpc.ClientConn) interface{} {
// 	return gaServices.NewCampaignServiceClient(conn)
// }

// func (s *CampaignService) GetCampaign(in *gaServices.GetCampaignRequest) error {
// 	campaign, error :=
// }

/*
	GoogleAdsService:
		- GoogleAdsServiceClient (interface)
			- Search()
		- SearchGoogleAdsRequest (struct)
		- SearchGoogleAdsResponse (struct)
		- GoogleAdsRow (struct)

	CampaignService:
		- CampaignServiceClient (interface)
			- GetCampaign
			- MutateCampaign
		- GetCampaignRequest (struct)
		- MutateCampaignRequest (struct)

*/
