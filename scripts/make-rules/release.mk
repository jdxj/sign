CHANGELOG_DIR := $(ROOT_DIR)/CHANGELOG

.PHONY: release.ensure-tag
release.ensure-tag:
	@scripts/ensure_tag.sh

.PHONY: release.tag
release.tag: release.ensure-tag
	@git push origin $(shell git describe --tags --abbrev=0)

.PHONY: release.chglog
release.chglog:
	@git-chglog -o $(CHANGELOG_DIR)/CHANGELOG-$(shell gsemver bump patch).md \
	$(shell git describe --tags --abbrev=0 --match 'v*')..