.PHONY: default clean_build clean_bin build_grpc build_grpc build_gopatrol build_gopatrol-cli

CLI_NAME = gopatrol-cli
DAEMON_NAME = gopatrol
OS := $(shell uname)
VERSION ?= 1.0.0

# test target

PROTOC_BIN=~/protoc/bin/protoc

# target #

default: clean_build build_gopatrol

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

generate_go: resources/* 
	./bin2go -in=resources/js/checkup.js -out=templates/checkupjs.go -pkg=templates CheckupJS
	./bin2go -in=resources/js/fs.js -out=templates/fsjs.go -pkg=templates FSJS
	./bin2go -in=resources/js/s3.js -out=templates/s3js.go -pkg=templates S3JS
	./bin2go -in=resources/js/statuspage.js -out=templates/statuspagejs.go -pkg=templates StatusPageJS
	./bin2go -in=resources/css/style.css -out=templates/css.go -pkg=templates StyleCSS
	./bin2go -in=resources/index.html -out=templates/html.go -pkg=templates IndexHTML
	./bin2go -in=resources/images/status-gray.png -out=templates/imagesgray.go -pkg=templates StatusGrayPNG
	./bin2go -in=resources/images/status-red.png -out=templates/imagesred.go -pkg=templates StatusRedPNG
	./bin2go -in=resources/images/status-yellow.png -out=templates/imagesyellow.go -pkg=templates StatusYellowPNG
	./bin2go -in=resources/images/status-green.png -out=templates/imagesgreen.go -pkg=templates StatusGreenPNG

clean_bin:
	- rm bin2go

clean_build:
	- rm -rf build/*
