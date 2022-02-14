TOOLS_DIR := $(ROOT_DIR)/tools

.PHONY: install.codegen
install.codegen:
	go install $(TOOLS_DIR)/codegen
