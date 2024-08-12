FROM --platform=linux/amd64 golang:1.22.5-alpine3.20 as builder

RUN apk add --no-cache build-base linux-headers gcc musl-dev git openssh-client

COPY ./keys/* /root/.ssh/

RUN chown -R root:root /root/.ssh
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts
RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"

WORKDIR /src

COPY go.mod .
COPY go.sum .

ENV GOPRIVATE="github.com/trylotus/*"

RUN go mod download

COPY . .

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=1
ENV GOFLAGS="-tags=musl"

RUN go mod tidy
RUN go build -o bin cmd/example/main.go

FROM --platform=linux/amd64 alpine:3.20

RUN apk add --no-cache musl

WORKDIR /app

COPY local.yaml .
COPY manifest.yaml .
COPY --from=builder /src/bin .

ENTRYPOINT [ "./bin" ]
