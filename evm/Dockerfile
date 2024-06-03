# FROM golang:1.21 as builder

# # RUN apk add alpine-sdk
# # RUN apk add --no-cache gcc musl-dev openssh-client

# RUN dpkg --add-architecture amd64 \
#     && apt-get update \
#     && apt-get install -y --no-install-recommends gcc-x86-64-linux-gnu libc6-dev-amd64-cross

# WORKDIR /connector

# COPY ./keys/* /root/.ssh/

# RUN chown -R root:root /root/.ssh
# RUN ssh-keyscan github.com >> /root/.ssh/known_hosts
# RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"

# COPY . .

# ENV GOPRIVATE="github.com/trylotus/*"

# RUN go mod tidy

# ENV GOOS=linux
# ENV GOARCH=amd64
# ENV CGO_ENABLED=1
# ENV GOFLAGS="-tags=musl"

# RUN go build -o app cmd/evm/main.go

# FROM alpine:3.20

# RUN apk add --no-cache musl

# WORKDIR /connector

# COPY local.yaml .
# COPY --from=builder app .

# ENTRYPOINT [ "./app" ]


FROM ubuntu:20.04 as builder

# Install the C lib for kafka
RUN apt-get update
RUN apt-get install -y --no-install-recommends apt-utils wget gnupg software-properties-common
RUN apt-get install -y apt-transport-https ca-certificates
RUN wget -qO - https://packages.confluent.io/deb/5.1/archive.key | apt-key add -
RUN add-apt-repository "deb [arch=amd64] https://packages.confluent.io/deb/5.1 stable main"
RUN apt-get update
RUN apt-get install -y librdkafka-dev

# Install Go
RUN add-apt-repository ppa:longsleep/golang-backports
RUN apt-get update
RUN apt-get install -y golang-1.21-go

# Build application
RUN apt-get -yqq install git ssh

COPY ./keys/* /root/.ssh/

RUN chown -R root:root /root/.ssh
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts
RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"

COPY . .

ENV GOPRIVATE="github.com/trylotus/*"

RUN go mod tidy

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=1
ENV GOFLAGS="-tags=musl"

RUN go build -o app cmd/evm/main.go

FROM ubuntu:20.04

WORKDIR /connector

COPY local.yaml .
COPY --from=builder app .

ENTRYPOINT [ "./app" ]
