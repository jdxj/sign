TOOLS_DIR := $(ROOT_DIR)/tools

.PHONY: install.codegen
install.codegen:
	@go install $(TOOLS_DIR)/codegen

.PHONY: install.signctl
install.signctl:
	@go install $(CMD_DIR)/signctl