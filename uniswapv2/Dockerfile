FROM --platform=linux/amd64 golang:1.23.0-alpine3.20 AS builder

RUN apk add --no-cache build-base linux-headers gcc musl-dev git openssh-client

COPY ./keys/* /root/.ssh/

RUN chown -R root:root /root/.ssh
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts
RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"

WORKDIR /src

COPY go.mod .
COPY go.sum .

ENV GOPRIVATE="github.com/trylotus/*"

RUN --mount=type=ssh go mod download

COPY . .

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=1
ENV GOFLAGS="-tags=musl"

RUN go mod tidy
RUN go build -o bin cmd/uniswapv2/main.go

FROM --platform=linux/amd64 alpine:3.20

RUN apk add --no-cache musl

WORKDIR /connector

ENV LOG_LEVEL=debug
ENV LOG_FORMAT=console
ENV KAFKA_BROKERS=localhost:9092
ENV REGISTRY_URL=http://localhost:8081

COPY manifest.yaml .
COPY --from=builder /src/bin .


ENTRYPOINT [ "./bin" ]
