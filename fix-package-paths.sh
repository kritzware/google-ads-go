MATCH="google\/ads\/googleads\/v0\/"
REPLACE="github.com\/kritzware\/google-ads-go\/protos\/google\/ads\/googleads\/v0\/"

PACKAGES=('common' 'enums' 'errors' 'resources' 'services')

echo "fixing package paths"
for file in ./protos/google/ads/googleads/v0/**/*.pb.go; do
    for p in "${PACKAGES[@]}"; do
        sed -i "" "s/$MATCH$p/$REPLACE$p/g" $file
    done
done
echo "finished fixing package paths"