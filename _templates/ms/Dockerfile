FROM golang:1.14.3-alpine3.11
WORKDIR /app
COPY . .
RUN apk update && apk add make && make build

FROM alpine:3.11
WORKDIR /app
COPY --from=0 /app/build ./
RUN chmod +x ./app-service
ENV ZEPTO_ENV=production
RUN apk update && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/* && \
    GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

ENTRYPOINT ["./app-service"]