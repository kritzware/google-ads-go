package main

import (
	"fmt"

	"github.com/kritzware/google-ads-go/ads"
	"github.com/kritzware/google-ads-go/services"
)

func main() {
	client, err := ads.NewClientFromStorage("google-ads.json")
	if err != nil {
		panic(err)
	}

	service := services.NewGoogleAdsServiceClient(client.Conn())

	request := services.SearchGoogleAdsRequest{
		CustomerId: "2984258132",
		Query:      "SELECT campaign.id, campaign.name FROM campaign ORDER BY campaign.id LIMIT 5",
	}

	response, err := service.Search(client.Context(), &request)

	for _, row := range response.Results {
		campaign := row.Campaign
		fmt.Printf("id: %d, name: %s\n", campaign.Id.Value, campaign.Name.Value)
	}
}
