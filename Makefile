GOPATH:=$(shell go env GOPATH)

.PHONY: update\:zepto
update\:zepto:
	go get -u github.com/go-zepto/zepto && \
	go mod tidy && \
	cd cmd/zepto/_templates/web && \
	mv go.mod_ go.mod && \
	go get -u github.com/go-zepto/zepto && \
	go mod tidy && \
	mv go.mod go.mod_ && \
	cd -

.PHONY: install
install:
	go get ./cmd/zepto
