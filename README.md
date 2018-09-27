# google-ads-go

**Note: This library is not ready for production yet**

| Google Ads API version 	| `v0` |
|-|:-:|
| Build | [![CircleCI](https://circleci.com/gh/kritzware/google-ads-go.svg?style=svg)](https://circleci.com/gh/kritzware/google-ads-go) |

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
