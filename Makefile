NAME=basht
VERSION=0.1.0

build:
	go-bindata include
	mkdir -p build/Linux  && GOOS=linux  go build -ldflags "-X main.Version $(VERSION)" -o build/Linux/$(NAME)
	mkdir -p build/Darwin && GOOS=darwin go build -ldflags "-X main.Version $(VERSION)" -o build/Darwin/$(NAME)

deps:
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/progrium/gh-release/...
	go get || true

test: build
	build/$(shell uname -s)/basht tests/meta.bash

.PHONY: build
