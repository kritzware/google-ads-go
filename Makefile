ADS_VERSION=v0
PROTO_ROOT_DIR=googleapis/
PROTO_SRC_DIR=/google/ads/googleads/$(ADS_VERSION)/**/*.proto
PROTO_OUT_DIR=$$GOPATH/src/github.com/kritzware/google-ads-go/
PKG_PATH=paths=source_relative
PROTOC_GO_ARGS=--go_out=plugins=grpc,$(PKG_PATH):$(PROTO_OUT_DIR)

ENTRY=main.go
BIN=gads

build:
	go build -o $$GOPATH/bin/$(BIN) $(ENTRY)

run:
	go run $(ENTRY)

run-debug:
	GODEBUG=http2debug=2 GRPC_GO_LOG_SEVERITY_LEVEL=info GRPC_GO_LOG_VERBOSITY_LEVEL=2 go run $(ENTRY)

test:
	go test ./

.SILENT protos: clean-protos clean-gen-protos
	echo "converting protos for version $(ADS_VERSION)"
	for file in $(PROTO_ROOT_DIR)$(PROTO_SRC_DIR); do \
		echo "converting proto $$(basename $$file)"; \
		protoc -I=$(PROTO_ROOT_DIR) $(PROTOC_GO_ARGS) $$file; \
	done; \
	sh ./fix-package-paths.sh; \
	rm -rf google/
	@echo "built proto files to $$(basename $(PROTO_OUT_DIR))"

clean-protos:
	rm -rf common/
	rm -rf enums/
	rm -rf errors/
	rm -rf resources/
	rm -rf services/

clean-gen-protos:
	rm -rf google/

.PHONY: protos
