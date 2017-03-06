.PHONY: clean
.PHONY: test

# Flags #
GO_FLAGS = -o

# path config #
PROTO_DIR=checkupservice/
PROTO_FILE=checkupservice.proto

# protoc
PROTOC_BIN=~/protoc/bin/protoc
PROTOC_OPTS=--go_out=plugins=grpc:checkupservice

# target #

default: clean build_protoc build_checkup

build_protoc:
	$(PROTOC_BIN) -I $(PROTO_DIR) $(PROTO_DIR)$(PROTO_FILE) $(PROTOC_OPTS)

build_checkup: 
	go build $(GO_FLAGS) checkup-server

clean:
	-rm checkup-server