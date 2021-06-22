

# protos:
# 	sh ./gapic.sh \
# 		--image gcr.io/gapic-images/gapic-generator-go \
# 		--in googleapis/google/ads/googleads/v8 \
# 		--out $(PWD)/build \
# 		--go-gapic-package "github.com/opteo/google-ads-go"

protos: build clean copy

build:
	DOCKER_BUILDKIT=1 docker build -t opteo/google-ads-go --progress=plain .
	
copy:
	docker create --name gads opteo/google-ads-go
	docker cp gads:/build $(PWD)
	docker rm gads
	chmod 700 $(PWD)/build
	cp -r build/google.golang.org .
	cp -r build/github.com/kritzware/google-ads-go/* gads/
	rm -rf build

clean:
	rm -rf build gads && mkdir build gads

.PHONY: build copy
	
# docker run opteo/google-ads-go

# # # --mount type=bind,source=</abs/path/to/configs>,destination=/conf,readonly \

# # protos:
# # 	docker run \
# # 		--rm \
# # 		--user $UID \
# # 		--mount type=bind,source=<./googleapis/google/ads/googleads/v8>,destination=/.,readonly \
# # 		--mount type=bind,source=$GOPATH/src,destination=/out/ \
# # 		gcr.io/gapic-images/gapic-generator-go \
# # 		--go-gapic-package "github.com/opteo/google-ads-go"