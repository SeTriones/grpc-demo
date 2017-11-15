DISTRIB_ROOT := $(shell pwd)
PROTO_DIR := $(DISTRIB_ROOT)/proto
CLIENT_DIR := $(DISTRIB_ROOT)/client
SERVER_DIR := $(DISTRIB_ROOT)/server
PROTO_FILES := $(shell find $(DISTRIB_ROOT) -name "*.proto")
PROTO_SROUCE := $(patsubst %.proto, %.pb.go, $(PROTO_FILES))

.PHONY: clean all

all: goproto client server

goproto: $(PROTO_SROUCE)

%.pb.go: %.proto
	cd $(PROTO_DIR) && protoc --go_out=plugins=grpc:. -I$(PROTO_DIR) $<

client: goproto $(CLIENT_DIR)/main.go
	cd $(CLIENT_DIR) && go build

server: goproto $(SERVER_DIR)/main.go 
	cd $(SERVER_DIR) && go build

clean:
	rm -rf ${DISTRIB_ROOT}/proto/*.pb.go
	rm -rf $(DISTRIB_ROOT)/client/client
	rm -rf $(DISTRIB_ROOT)/server/server
