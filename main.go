package main

import (
	"fmt"

	gads "github.com/kritzware/google-ads-go/client"
)

func main() {
	client, err := gads.NewGoogleAdsClientFromStorage("google-ads.json")
	if err != nil {
		panic(err)
	}

	campaignService := client.NewCampaignService()
	adGroupService := client.NewAdGroupService()

	campaign, err := campaignService.GetCampaign("3827277046", "954375723")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name: %s, Status: %s\n", campaign.Name.Value, campaign.ServingStatus)

	adGroup, err := adGroupService.GetAdGroup("3827277046", "51149598601")
	if err != nil {
		panic(err)
	}
	fmt.Println(adGroup)

	// service := client.GetService("CampaignService")
	// fmt.Printf("%+v\n", service)

	// campaignService := services.NewCampaignService(client.Conn())
	// fmt.Printf("service: %+v\n", campaignService)

	// campaign := campaignService.GetCampaign()
	// fmt.Printf("campaign: %+v\n", campaign)

	// load via string
	// service := client.NewGoogleAdsService(&client)
	// service := client.GetService("CampaignService")
	// fmt.Printf("%+v\n", service)

	// assign via pointer
	// var campaignService gads.Service
	// client.LoadService("CampaignService", &campaignService)
	// fmt.Printf("service: %+v\n", campaignService)

	// campaignService := services.NewCampaignService(&client)
	// adGroupAdService := services.NewAdGroupAdService(&client)
	// googleAdsService := services.NewGoogleAdsService(&client)

	// service := client.GetService("GoogleAdsService")
	// fmt.Printf("%+v\n", service)

	// request := client.SearchGoogleAdsRequest{
	// 	CustomerId: "3827277046",
	// 	Query:      "SELECT campaign.id, campaign.name FROM campaign ORDER BY campaign.id",
	// }
	// fmt.Printf("%+v\n", request)

	// response, err := service.Search(context.Background(), &request)
	// _, err = ga.Search(ctx, &request)
	// if err != nil {
	// 	s := status.Convert(err)
	// 	log.Printf("Response status: %s (code: %s)", s.Message(), s.Code())
	// }
	// fmt.Printf("%+v\n", service)
}
