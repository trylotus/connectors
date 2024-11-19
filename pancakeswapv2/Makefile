.PHONY: default clean gen_abi gen_proto

default: install_tools gen_abi gen_proto

install_tools:
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/bufbuild/buf/cmd/buf@latest

gen-abi:
	abigen --abi=./bep20/bep20.abi.json --pkg=bep20 --out=./bep20/bep20.abi.go
	abigen --abi=./factory/factory.abi.json --pkg=factory --out=./factory/factory.abi.go
	abigen --abi=./pair/pair.abi.json --pkg=pair --out=./pair/pair.abi.go

gen-proto:
	buf lint
	buf format -w
	buf generate
