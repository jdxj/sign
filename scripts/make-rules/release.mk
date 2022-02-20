CHANGELOG_DIR := $(ROOT_DIR)/CHANGELOG

.PHONY: release.ensure-tag
release.ensure-tag:
	@scripts/ensure_tag.sh

.PHONY: release.tag
release.tag: release.ensure-tag
	@git push origin $(shell git describe --tags --abbrev=0 --match 'v*')

.PHONY: release.chglog
release.chglog: release.ensure-tag
	$(eval VERSION := $(shell git describe --tags --abbrev=0 --match 'v*'))
	$(eval VERSION_NUM := $(shell echo $(VERSION) | sed 's/v//g'))
	@git-chglog -o $(CHANGELOG_DIR)/CHANGELOG-$(VERSION_NUM).md $(VERSION)