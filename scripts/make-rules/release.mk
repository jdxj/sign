.PHONY: release.ensure-tag
release.ensure-tag:
	scripts/ensure_tag.sh

.PHONY: release.tag
release.tag: release.ensure-tag
	git push origin $(shell git describe --tags --abbrev=0)