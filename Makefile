include scripts/make-rules/common.mk
include scripts/make-rules/gen.mk
include scripts/make-rules/go.mk
include scripts/make-rules/image.mk
include scripts/make-rules/release.mk
include scripts/make-rules/tools.mk

define USAGE_OPTIONS
Options:
  clean

  gen.code
  gen.proto.%
  gen.proto

  go.build.%
  go.build
  go.lint

  image.build.%
  image.build
  image.push.%
  image.push
  image.test-tag

  release.ensure-tag
  release.tag
  release.chglog

  install.codegen
endef
export USAGE_OPTIONS

.PHONY: help
help:
	@echo "$$USAGE_OPTIONS"
