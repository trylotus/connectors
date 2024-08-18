.PHONY: default clean gen_abi gen_proto

default: install_tools gen_abi gen_proto

install_tools:
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/bufbuild/buf/cmd/buf@latest

gen_proto:
	buf dep update
	buf build
	buf generate

gen_abi:
	abigen --abi=./xo/xo.abi --pkg=xo --out=./xo/xo.abi.go
