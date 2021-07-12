#!/bin/bash

VERSION=$1

protoc -I=/opt/include \
    --go_out=plugins=grpc:/build \
    --go_gapic_opt 'go-gapic-package=github.com/opteo/google-ads-go;gads' \
    --go_gapic_opt 'api-service-config=/opt/include/google/ads/googleads/${VERSION}/googleads_${VERSION}.yaml' \
    --go_gapic_opt 'grpc-service-config=/opt/include/google/ads/googleads/${VERSION}/googleads_grpc_service_config.json' \
    --go_gapic_out=/build \
    /opt/include/google/ads/googleads/${VERSION}/**/**.proto
