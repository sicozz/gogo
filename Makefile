APP				?= gogo
APP_VERSION		?= 0.0.0
SERVER_ENTRY	?= server/main.go
SERVER_BIN		?= bin/$(APP)
API_VERSION		?= 0
PROTO_DIR		?= api/v$(API_VERSION)

.PHONY: all
all: build

# Build
.PHONY: build
build: build-proto
	@ go build -o $(SERVER_BIN) $(SERVER_ENTRY)

.PHONY: build-proto
build-proto:
	@ protoc --proto_path=$(PROTO_DIR) \
       --go_out=$(PROTO_DIR) --go_opt=paths=source_relative \
       --go-grpc_out=$(PROTO_DIR) --go-grpc_opt=paths=source_relative \
       $(PROTO_DIR)/$(APP).proto

.PHONY: debug
debug:
	dlv debug $(SERVER_ENTRY)

# Cleanup
.PHONY: clean
clean:
	rm bin/*

# Dev dependencies
.PHONY: install-dev-deps
install-dev-deps:
	@ go install github.com/go-delve/delve/cmd/dlv@latest
	@ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

