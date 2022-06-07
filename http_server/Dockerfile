FROM golang:1.16 AS builder

COPY . /root

WORKDIR /root

RUN mkdir -p bin/ && GOPROXY=https://goproxy.cn go build -o ./bin/httpServer main.go

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends

COPY --from=builder /root/bin /httpServer

WORKDIR /httpServer

EXPOSE 80

ENTRYPOINT ["./httpServer"]