FROM node:16-alpine3.14
WORKDIR /app
# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
# Copy Zepto Project
COPY . .
## Install dependencies and build
RUN apk update && \
    apk add --update go=1.16.5-r0 git alpine-sdk make python3 && \
    go get -u github.com/go-zepto/zepto-cli/cmd/zepto && \
    export PATH=$PATH:/usr/local/go/bin && \
    npm install && \
    go mod tidy && \
    zepto build

FROM alpine:3.14
WORKDIR /app
COPY --from=0 /app/build ./
RUN chmod +x ./app
ENV ZEPTO_ENV=production
RUN apk update && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

ENTRYPOINT ["./app"]
