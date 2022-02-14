include scripts/make-rules/common.mk
include scripts/make-rules/gen.mk
include scripts/make-rules/go.mk
include scripts/make-rules/image.mk
include scripts/make-rules/release.mk
include scripts/make-rules/tools.mk

.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	rm -rf build/output
