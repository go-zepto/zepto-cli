GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc --proto_path=. --go_out=plugins=grpc:. proto/app/app.proto


.PHONY: build
build:
	export ZEPTO_ENV=production && \
	rm -rf public/build && \
	rm -rf build && mkdir build && \
	go build -o build/app-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t app-service:latest
