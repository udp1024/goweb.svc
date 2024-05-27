#!/bin/bash

platforms=("windows/amd64" "windows/arm64" "linux/amd64" "linux/arm64" "darwin/arm64" "darwin/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name='bin/goweb-svc-'$GOOS'-'$GOARCH
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name .
    echo 'Built for '$platform
done
