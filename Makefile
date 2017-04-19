.PHONY: default clean

CLI_NAME = checklist
DAEMON_NAME = checkupd
OS := $(shell uname)
VERSION ?= 1.0.0

# test target

PROTOC_BIN=~/protoc/bin/protoc

# target #

default: clean build_checkupd build_checklist

build_grpc:
	$(PROTOC_BIN) -I service/ --go_out=plugins=grpc:service grpc/service/service.proto

build_docker: 
	docker build --tag checkupd:v${VERSION} .

build_checkupd: 
	@echo "Setup checkupd"
ifeq ($(OS),Linux)
	mkdir -p build/linux
	@echo "Build checkupd..."
	GOOS=linux  go build -ldflags "-s -w -X main.Version=$(VERSION)" -o build/linux/$(DAEMON_NAME) cmd/server/main.go
	sudo cp ./build/linux/checkupd /usr/local/bin/
endif
ifeq ($(OS) ,Darwin)
	@echo "Build checkupd..."
	GOOS=darwin go build -ldflags "-X main.Version=$(VERSION)" -o build/mac/$(DAEMON_NAME) cmd/server/main.go
	sudo cp ./build/mac/checkupd /usr/local/bin/
endif
	@echo "Succesfully Build for ${OS} version:= ${VERSION}"

build_checklist: 
	@echo "Setup checklist"
ifeq ($(OS),Linux)
	mkdir -p build/linux
	@echo "Build checklist..."
	GOOS=linux  go build -ldflags "-s -w -X main.Version=$(VERSION)" -o build/linux/$(CLI_NAME) cmd/cli/main.go
	sudo cp ./build/linux/checklist /usr/local/bin/
endif
ifeq ($(OS) ,Darwin)
	@echo "Build checklist..."
	GOOS=darwin go build -ldflags "-X main.Version=$(VERSION)" -o build/mac/$(CLI_NAME) cmd/cli/main.go
	sudo cp ./build/mac/checklist /usr/local/bin/
endif
	@echo "Succesfully Build for ${OS} version:= ${VERSION}"

clean:
	rm -rf build/*

reset_setup:
	- rm -rf checkup_config
	- rm -rf caddy_config
