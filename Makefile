.PHONY: default clean

CLI_NAME = checklist
DAEMON_NAME = checkupd
OS := $(shell uname)
VERSION ?= 1.0.0

# test target

# target #

default: clean build_checkupd build_checklist

build_docker: 
	docker build --tag checkupd:v${VERSION} .

clean_logs:
	- rm -rf caddy-errors
	- rm -rf caddy-logs
	- rm -rf logs

build_checkupd: 
	@echo "Setup checkupd"
ifeq ($(OS),Linux)
	mkdir -p build/linux
	@echo "Build checkupd..."
	GOOS=linux  go build -ldflags "-s -w -X main.Version=$(VERSION)" -o build/linux/$(DAEMON_NAME) cmd/server/main.go
endif
ifeq ($(OS) ,Darwin)
	@echo "Build checkupd..."
	GOOS=darwin go build -ldflags "-X main.Version=$(VERSION)" -o build/mac/$(DAEMON_NAME) cmd/server/main.go
endif
	@echo "Succesfully Build for ${OS} version:= ${VERSION}"

build_checklist: 
	@echo "Setup checklist"
ifeq ($(OS),Linux)
	mkdir -p build/linux
	@echo "Build checklist..."
	GOOS=linux  go build -ldflags "-s -w -X main.Version=$(VERSION)" -o build/linux/$(CLI_NAME) cmd/cli/main.go
endif
ifeq ($(OS) ,Darwin)
	@echo "Build checklist..."
	GOOS=darwin go build -ldflags "-X main.Version=$(VERSION)" -o build/mac/$(CLI_NAME) cmd/cli/main.go
endif
	@echo "Succesfully Build for ${OS} version:= ${VERSION}"

clean:
	rm -rf build/*