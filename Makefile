GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	pkger -include github.com/go-zepto/zepto-cli:/_templates