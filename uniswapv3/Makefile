.PHONY: default clean gen_abi gen_proto

default: install_tools gen_abi gen_proto

install_tools:
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/bufbuild/buf/cmd/buf@latest

gen-proto:
	buf lint
	buf format -w
	buf generate

gen-abi:
	abigen --abi=./erc20/erc20.abi --pkg=erc20 --out=./erc20/erc20.abi.go
	abigen --abi=./factory/factory.abi --pkg=factory --out=./factory/factory.abi.go
	abigen --abi=./pool/pool.abi --pkg=pool --out=./pool/pool.abi.go