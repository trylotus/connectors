FROM blepai/golang-proto:latest AS builder

RUN apk add --no-cache build-base linux-headers

# download dependencies first
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# build protobufs
RUN protoc --proto_path=. \
                --go_out=. --go_opt=paths=source_relative \
                --go-grpc_out=. --go-grpc_opt=paths=source_relative --experimental_allow_proto3_optional **/*.proto

# build. CGO_ENABLED=1 required for conflient-kafka-go
RUN mkdir /src/bin
RUN CGO_ENABLED=1 GOFLAGS="-tags=musl" go build -o bin ./...

#FROM scratch AS final
# can't use scratch for confluent-kafka-go
FROM alpine:3.15 AS final

COPY --from=builder /src/bin /usr/local/bin
COPY --from=builder /src/manifest.yaml /

CMD ["ash"]
