.PHONY: default clean_build clean_bin build_grpc build_grpc build_gopatrol build_gopatrol-cli

CLI_NAME = gopatrol-cli
DAEMON_NAME = gopatrol
OS := $(shell uname)
VERSION ?= 1.0.0

PKG_NAME = github.com/AdhityaRamadhanus/gopatrol
TEST_PKG = ${PKG_NAME} ${PKG_NAME}/api ${PKG_NAME}/api/helper

# test target

PROTOC_BIN=~/protoc/bin/protoc

# target #

default: clean_build build_gopatrol build_gopatrol_cli

build_gopatrol: 
	@echo "Setup gopatrol"
ifeq ($(OS),Linux)
	mkdir -p build/linux
	@echo "Build gopatrol..."
	GOOS=linux  go build -ldflags "-s -w -X main.Version=$(VERSION)" -o build/linux/$(DAEMON_NAME) cmd/server/main.go
	sudo cp ./build/linux/gopatrol /usr/local/bin/
endif
ifeq ($(OS) ,Darwin)
	@echo "Build gopatrol..."
	GOOS=darwin go build -ldflags "-X main.Version=$(VERSION)" -o build/mac/$(DAEMON_NAME) cmd/server/main.go
	sudo cp ./build/mac/gopatrol /usr/local/bin/
endif
	@echo "Succesfully Build for ${OS} version:= ${VERSION}"

build_gopatrol_cli: 
	@echo "Setup gopatrol cli"
ifeq ($(OS),Linux)
	mkdir -p build/linux
	@echo "Build gopatrol-cli..."
	GOOS=linux  go build -ldflags "-s -w -X main.Version=$(VERSION)" -o build/linux/$(CLI_NAME) cmd/cli/main.go
	sudo cp ./build/linux/gopatrol-cli /usr/local/bin/
endif
ifeq ($(OS) ,Darwin)
	@echo "Build gopatrol..."
	GOOS=darwin go build -ldflags "-X main.Version=$(VERSION)" -o build/mac/$(CLI_NAME) cmd/server/main.go
	sudo cp ./build/mac/gopatrol-cli /usr/local/bin/
endif
	@echo "Succesfully Build for ${OS} version:= ${VERSION}"


# Test Packages

test:
	go test -v --cover ${TEST_PKG}

clean_build:
	- rm -rf build/*
