#!/bin/bash

protoc -I=/opt/include \
    --go_out=plugins=grpc:/build \
    --go_gapic_opt 'go-gapic-package=github.com/kritzware/google-ads-go;gads' \
    --go_gapic_opt 'api-service-config=/opt/include/google/ads/googleads/v8/googleads_v8.yaml' \
    --go_gapic_opt 'grpc-service-config=/opt/include/google/ads/googleads/v8/googleads_grpc_service_config.json' \
    --go_gapic_out=/build \
    /opt/include/google/ads/googleads/v8/**/**.proto

# PROTOS_ENUMS="/go/googleapis/google/ads/googleads/v8/enums/*.proto"
# PROTOS_ERRORS="/go/googleapis/google/ads/googleads/v8/errors/*.proto"
# PROTOS_COMMON="/go/googleapis/google/ads/googleads/v8/common/*.proto"
# PROTOS_RESOURCES="/go/googleapis/google/ads/googleads/v8/resources/*.proto"
# PROTOS_SERVICES="/go/googleapis/google/ads/googleads/v8/services/*.proto"

# protoc --version

# protoc -I=/go/googleapis \
#     --go_gapic_opt 'go-gapic-package=github.com/opteo/google-ads-go;google-ads-go' \
#     --go_gapic_opt 'api-service-config=/go/googleapis/google/ads/googleads/v8/googleads_v8.yaml' \
#     --go_gapic_opt 'grpc-service-config=/go/googleapis/google/ads/googleads/v8/googleads_grpc_service_config.json' \
#     --go_gapic_out=/build \
#     /go/googleapis/google/ads/googleads/v8/**/**.proto

# function compile_protos() {
#     PROTOS=$*
#     for file in $PROTOS; do
#         echo "converting proto $(basename $file)"
#         # FOLDER_PATH=$(dirname $file)
#         # PACKAGE=$(basename $FOLDER_PATH)
#         # GO_PACKAGE="\"github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/$PACKAGE\""
#         # echo "option go_package = $GO_PACKAGE;" >> $file
#         protoc -I=/go/googleapis \
#             --go_gapic_opt 'go-gapic-package=github.com/opteo/google-ads-go;google-ads-go' \
#             --go_gapic_opt 'api-service-config=/go/googleapis/google/ads/googleads/v8/googleads_v8.yaml' \
#             --go_gapic_opt 'grpc-service-config=/go/googleapis/google/ads/googleads/v8/googleads_grpc_service_config.json' \
#             --go_gapic_out=/build \
#             $file
#     done
# }

# compile_protos $PROTOS_ENUMS
# # compile_protos $PROTOS_ERRORS
# # compile_protos $PROTOS_COMMON
# # compile_protos $PROTOS_RESOURCES
# # compile_protos $PROTOS_SERVICES

# ls -l -a /build

# # RUN 

# protoc \
#     -I /go/googleapis \
#     --go_gapic_out /go/ \
#     --go_gapic_opt 'go-gapic-package=github.com/opteo/google-ads-go;google-ads-go' \
#     --go_gapic_out 'api-service-config=./google/ads/googleads/v8/googleads_v8.yaml' \
#     --go_gapic_out 'grpc-service-config=./google/ads/googleads/v8/googleads_grpc_service_config.json' \
#     --proto_path=/go/googleapis --go_out=out --go_opt=paths=source_relative \
#     /go/googleapis/google/ads/googleads/v8/**/**.proto

# ls -l -a /go/