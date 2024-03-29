DOCKER_DIR := $(BUILD_DIR)/docker
TMP_DIR := $(BUILD_DIR)/tmp
REGISTRY_PREFIX := jdxj
VERSION := $(shell git describe --tags --abbrev=0)
DOCKERS := $(filter-out %.sh, $(wildcard $(DOCKER_DIR)/*))
CONTAINERS := $(foreach doc, $(DOCKERS), $(notdir $(doc)))

.PHONY: image.build.%
image.build.%: go.build.%
	@mkdir -p $(TMP_DIR)/$*
	@cat $(DOCKER_DIR)/$*/Dockerfile > $(TMP_DIR)/$*/Dockerfile
	@cp $(GO_BUILD_DIR)/$*$(GO_EXT) $(TMP_DIR)/$*
	@docker build -t $(REGISTRY_PREFIX)/$*:$(VERSION) $(TMP_DIR)/$*
	@rm -rf $(TMP_DIR)/$*

.PHONY: image.build
image.build: $(addprefix image.build., $(CONTAINERS))

.PHONY: image.push.%
image.push.%: image.build.%
	docker push $(REGISTRY_PREFIX)/$*:$(VERSION)

.PHONY: image.push
image.push: $(addprefix image.push., $(CONTAINERS))

.PHONY: image.test-tag
image.test-tag:
	@scripts/test_tag.sh
