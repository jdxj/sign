#!/usr/bin/env bash

if [ -z "$1" ]; then
  echo "module is required"
  exit 1
fi

echo "build $1 image"

workDir=$(pwd)
commit=$(git rev-parse --short HEAD)
tag=jdxj/$1:test-$commit

cd "$workDir/$1" && docker build -t="$tag" .
if [ $? -ne 0 ]; then
  echo "build $1 image failed"
  exit $?
fi

docker push "$tag"
