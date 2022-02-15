#!/usr/bin/env bash

version=test-$(git rev-parse --short HEAD)
if [ -z "$(git tag -l $version)" ]; then
  git tag -a -m "test version $version" $version
fi
