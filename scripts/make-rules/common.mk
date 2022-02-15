ROOT_DIR := .
BUILD_DIR := $(ROOT_DIR)/build
CMD_DIR := $(ROOT_DIR)/cmd

.PHONY: clean
clean:
	@rm -rf $(BUILD_DIR)/output
	@rm -rf $(BUILD_DIR)/tmp
