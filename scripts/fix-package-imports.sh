PACKAGES=('common' 'enums' 'errors' 'resources' 'services')
VERSION=$1

function fix_package_path() {
    FILE=$1
    PACKAGE=$2
    MATCH="google.golang.org\/genproto\/googleapis\/ads\/googleads\/${VERSION}\/"
    REPLACE="github.com\/opteo\/google-ads-go\/"
    sed -i "" "s/$MATCH$PACKAGE/$REPLACE$PACKAGE/g" $FILE
}

function fix_package_name() {
    FILE=$1
    PACKAGE=$2
    sed -i "" "s/google_ads_googleads_${VERSION}_$PACKAGE/$PACKAGE/g" $FILE
}

echo "fixing package imports.."
for file in $(PWD)/google.golang.org/genproto/googleapis/ads/googleads/${VERSION}/**/*.pb.go; do
    for p in "${PACKAGES[@]}"; do
        fix_package_path $file $p
        fix_package_name $file $p
    done
done

echo "fixing client imports.."
for file in $(PWD)/gads/*.go; do
    for p in "${PACKAGES[@]}"; do
        fix_package_path $file $p
        fix_package_name $file $p
    done
done

echo "moving compiled packages to root directory.."
for p in "${PACKAGES[@]}"; do
    mv "$(PWD)/google.golang.org/genproto/googleapis/ads/googleads/${VERSION}/$p" $(PWD)/
done

rm -rf ./google.golang.org
echo "all done"