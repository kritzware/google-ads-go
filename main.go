package main

import (
	"fmt"

	"github.com/kritzware/google-ads-go/client"
)

func main() {
	client := client.NewGoogleAdsClient(&client.GoogleAdsClientArgs{
		ClientID:       "",
		ClientSecret:   "",
		DeveloperToken: "",
		RefreshToken:   "",
	})

	service := client.GetService("GoogleAdsService")

	fmt.Printf("%+v\n", service)

	// request := services.SearchGoogleAdsRequest{
	// 	CustomerId: "3827277046",
	// 	Query:      "SELECT campaign.id, campaign.name FROM campaign ORDER BY campaign.id",
	// }

	// response, err := service.Search(context.Background(), &request)
	// _, err = ga.Search(ctx, &request)
	// if err != nil {
	// 	s := status.Convert(err)
	// 	log.Printf("Response status: %s (code: %s)", s.Message(), s.Code())
	// }
	// fmt.Printf("%+v\n", service)
}
