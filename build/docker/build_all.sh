#!/usr/bin/env bash

workDir=$(pwd)
version=$(git describe --tags --abbrev=0)
modules=$(ls -l | awk '/^d/ {print $NF}')

for module in $modules; do
  echo "build $module image"

  tag=jdxj/$module:$version
  cd "$workDir/$module" && docker build -t="$tag" .
  if [ $? -ne 0 ]; then
    echo "build $module image failed"
    exit $?
  fi
  docker push "$tag"

  echo
done
