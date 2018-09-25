# google-ads-go

**Note: This library is not ready for production yet**

| Google Ads API version 	| `v0` |
|-|:-:|
| Build | [![CircleCI](https://circleci.com/gh/kritzware/google-ads-go.svg?style=svg)](https://circleci.com/gh/kritzware/google-ads-go) |

## Installation
To install, simply run:
```bash
$ go get github.com/kritzware/google-ads-go
```
Make sure your PATH includes the $GOPATH/bin directory if you want to use the CLI utils:
```bash
export PATH=$PATH:$GOPATH/bin
````

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
