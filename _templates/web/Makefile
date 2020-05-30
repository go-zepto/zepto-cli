GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:
	export NODE_ENV=production && \
	export ZEPTO_ENV=production && \
	rm -rf public/build && \
	npm run build && \
	rm -rf build && mkdir build && \
	go build -o build/app-service *.go &&\
	cp -r templates public ./build

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t app-service:latest
