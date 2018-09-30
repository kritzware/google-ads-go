package ads

import "github.com/kritzware/google-ads-go/client"

func NewClientFromStorage(filepath string) (*client.GoogleAdsClient, error) {
	return client.NewGoogleAdsClientFromStorage(filepath)
}

func NewClient(params *client.GoogleAdsClientParams) (*client.GoogleAdsClient, error) {
	return client.NewGoogleAdsClient(params)
}
