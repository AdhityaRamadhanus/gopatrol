.PHONY: default clean

CLI_NAME = gopatrol-cli
DAEMON_NAME = gopatrol
OS := $(shell uname)
VERSION ?= 1.0.0

# test target

PROTOC_BIN=~/protoc/bin/protoc

# target #

default: clean build_gopatrol build_gopatrol-cli

build_grpc:
	$(PROTOC_BIN) -I service/ --go_out=plugins=grpc:service grpc/service/service.proto

build_docker: 
	docker build --tag gopatrol:v${VERSION} .

build_gopatrol: 
	@echo "Setup gopatrol"
ifeq ($(OS),Linux)
	mkdir -p build/linux
	@echo "Build gopatrol..."
	GOOS=linux  go build -ldflags "-s -w -X main.Version=$(VERSION)" -o build/linux/$(DAEMON_NAME) cmd/daemon/main.go
	sudo cp ./build/linux/gopatrol /usr/local/bin/
endif
ifeq ($(OS) ,Darwin)
	@echo "Build gopatrol..."
	GOOS=darwin go build -ldflags "-X main.Version=$(VERSION)" -o build/mac/$(DAEMON_NAME) cmd/daemon/main.go
	sudo cp ./build/mac/gopatrol /usr/local/bin/
endif
	@echo "Succesfully Build for ${OS} version:= ${VERSION}"

build_gopatrol-cli: 
	@echo "Setup gopatrol-cli"
ifeq ($(OS),Linux)
	mkdir -p build/linux
	@echo "Build gopatrol-cli..."
	GOOS=linux  go build -ldflags "-s -w -X main.Version=$(VERSION)" -o build/linux/$(CLI_NAME) cmd/cli/main.go
	sudo cp ./build/linux/gopatrol-cli /usr/local/bin/
endif
ifeq ($(OS) ,Darwin)
	@echo "Build gopatrol-cli..."
	GOOS=darwin go build -ldflags "-X main.Version=$(VERSION)" -o build/mac/$(CLI_NAME) cmd/cli/main.go
	sudo cp ./build/mac/gopatrol-cli /usr/local/bin/
endif
	@echo "Succesfully Build for ${OS} version:= ${VERSION}"

clean:
	rm -rf build/*

reset_setup:
	- rm -rf checkup_config
	- rm -rf caddy_config
