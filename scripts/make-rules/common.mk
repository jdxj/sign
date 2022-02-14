ROOT_DIR := .
BUILD_DIR := $(ROOT_DIR)/build
CMD_DIR := $(ROOT_DIR)/cmd

.PHONY: root
root:
	@echo $(ROOT_DIR)
	@echo $(BUILD_DIR)