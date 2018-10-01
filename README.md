# google-ads-go

| Google Ads API version 	| `v0` |
|-|:-:|
| Build | [![CircleCI](https://circleci.com/gh/kritzware/google-ads-go.svg?style=svg)](https://circleci.com/gh/kritzware/google-ads-go) |

## Features
- Fully matches the latest [Google Ads API Reference](https://developers.google.com/google-ads/api/reference/rpc/)
- Implemented via [gRPC](https://grpc.io/) with [Protocol Buffers](https://developers.google.com/protocol-buffers/)
- CLI utils and API for generating and refreshing Google Ads credentials

## Installation
To install, simply run:
```bash
$ go get -d github.com/kritzware/google-ads-go
```
Make sure your PATH includes the $GOPATH/bin directory if you want to use the CLI utils:
```bash
export PATH=$PATH:$GOPATH/bin
````

## Example
```go
package main

import (
  "fmt"

  "github.com/kritzware/google-ads-go/ads"
  "github.com/kritzware/google-ads-go/services"
)

func main() {
  // Create a client from credentials file
  client, err := ads.NewClientFromStorage("google-ads.json")
  if err != nil {
    panic(err)
  }
  
  // Load the "GoogleAds" service
  googleAdsService := services.NewGoogleAdsServiceClient(client.Conn())
  
  // Create a search request
  request := services.SearchGoogleAdsRequest{
    CustomerId: "2984242032",
    Query:      "SELECT campaign.id, campaign.name FROM campaign ORDER BY campaign.id",
  }
  
  // Get the results
  response, err := googleAdsService.Search(client.Context(), &request)
  for _, row := range response.Results {
    campaign := row.Campaign
    fmt.Printf("id: %d, name: %s\n", campaign.Id.Value, campaign.Name.Value)
  }
}
```

When using the `NewGoogleAdsClientFromStorage` method, you must provide a path to a valid `google-ads.json` file (containing your Google Ads API credentials).
```json
{
    "client_id": "YOUR_CLIENT_ID",
    "client_secret": "YOUR_CLIENT_SECRET",
    "refresh_token": "YOUR_REFRESH_TOKEN",
    "developer_token": "YOUR_DEVELOPER_TOKEN"
}

```

## Contributing
### Compiling
All build scripts use `Makefile`

**Build and run**
```bash
$ make run
```

**Run with gRPC debugging output**
```bash
$ make run-debug
```

### Manually Building Protos
Building `.pb.go` files from the original `googleads` protos should only be done when updating to a new Google Ads version.

Requirements:
- [protoc](https://github.com/protocolbuffers/protobuf)

Build `.pb.go` protos:
```bash
$ make protos
```
