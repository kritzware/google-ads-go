# Args from make
PROTO_ROOT_DIR=$1
PROTO_SRC_DIR=$2
PROTOC_GO_ARGS=$3

# Protos in order to compile
PROTOS_ENUMS=googleapis/google/ads/googleads/v0/enums/*.proto
PROTOS_ERRORS=googleapis/google/ads/googleads/v0/errors/*.proto
PROTOS_COMMON=googleapis/google/ads/googleads/v0/common/*.proto
PROTOS_RESOURCES=googleapis/google/ads/googleads/v0/resources/*.proto
PROTOS_SERVICES=googleapis/google/ads/googleads/v0/services/*.proto

function compile_protos() {
    PROTOS=$*
    for file in $PROTOS; do
        echo "converting proto $(basename $file)"
        FOLDER_PATH=$(dirname $file)
        PACKAGE=$(basename $FOLDER_PATH)
        # GO_PACKAGE="\"github.com/kritzware/google-ads-go/protos/google/ads/googleads/v0/$PACKAGE\""
        # echo "option go_package = $GO_PACKAGE;" >> $file
        protoc -I=$PROTO_ROOT_DIR $PROTOC_GO_ARGS $file
    done
}

compile_protos $PROTOS_ENUMS
compile_protos $PROTOS_ERRORS
compile_protos $PROTOS_COMMON
compile_protos $PROTOS_RESOURCES
compile_protos $PROTOS_SERVICES

cd googleapis && git checkout . && cd ../