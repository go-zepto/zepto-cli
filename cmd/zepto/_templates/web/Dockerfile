FROM golang:1.16.6-alpine3.14
WORKDIR /app
COPY . .
RUN apk update && \
    apk add --update git alpine-sdk make npm python3 && \
    go get -u github.com/go-zepto/zepto-cli/cmd/zepto && \
    npm install && zepto build

FROM alpine:3.11
WORKDIR /app
COPY --from=0 /app/build ./
RUN chmod +x ./app
ENV ZEPTO_ENV=production
RUN apk update && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

ENTRYPOINT ["./app"]
