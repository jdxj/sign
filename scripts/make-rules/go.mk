GO_BUILD_DIR := $(BUILD_DIR)/output
GO_EXT := .out
COMMANDS := $(filter-out %test, $(wildcard $(CMD_DIR)/*))
BINS := $(foreach cmd, $(COMMANDS), $(notdir $(cmd)))

.PHONY: hello
hello:
	@echo $(GO_BUILD_DIR)
	@echo $(COMMANDS)
	@echo $(BINS)

.PHONY: go.build.%
go.build.%:
	@mkdir -p $(GO_BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $(GO_BUILD_DIR)/$*$(GO_EXT) $(CMD_DIR)/$*/*.go

.PHONY: go.build
go.build: $(addprefix go.build., $(BINS))