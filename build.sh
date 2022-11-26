#!/bin/bash

dist_dir="dist"

build_oss=("darwin" "linux" "windows")
build_archs=("amd64" "arm64")


for GOOS in ${build_oss[@]}; do
  for GOARCH in ${build_archs[@]}; do
    echo "Building for $GOOS/$GOARCH"
    echo
    env GOOS=$GOOS GOARCH=$GOARCH go build -o "$dist_dir/go-passphrase_$GOOS-$GOARCH" github.com/tashima42/go-passphrase 
    echo "====================="
  done
done