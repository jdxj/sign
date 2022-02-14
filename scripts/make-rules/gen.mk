PROTO_DIR := $(ROOT_DIR)/internal/proto
PROTOS := $(filter-out %test-grpc, $(wildcard $(PROTO_DIR)/*))
PBS := $(foreach pb, $(PROTOS), $(notdir $(pb)))

.PHONY: gen.code
gen.code:
	go generate ./...

.PHONY: gen.proto.%
gen.proto.%:
	@protoc --proto_path=.:/usr/local/include \
	--go_out=. --go_opt=paths=source_relative \
	--micro_out=. --micro_opt=paths=source_relative \
	$(PROTO_DIR)/$*/*.proto

.PHONY: gen.proto
gen.proto: $(addprefix gen.proto., $(PBS))
