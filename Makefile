GOOGLE_ADS_VERSION=v8

protos: build clean copy fix-imports

build:
	DOCKER_BUILDKIT=1 docker build \
		--build-arg GOOGLE_ADS_VERSION=${GOOGLE_ADS_VERSION} \
		--progress=plain \
		-t opteo/google-ads-go .
	
copy:
	docker create --name gads opteo/google-ads-go
	docker cp gads:/build $(PWD)
	docker rm gads
	chmod 700 $(PWD)/build
	cp -r build/google.golang.org $(PWD)
	cp -r build/github.com/opteo/google-ads-go/* gads/
	rm -rf build

clean:
	rm -rf build gads && mkdir build gads

fix-imports:
	sh ./scripts/fix-package-imports.sh ${GOOGLE_ADS_VERSION}
	go mod tidy

.PHONY: protos build copy clean fix-imports