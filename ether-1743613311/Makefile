.PHONY: default clean gen-abi gen-proto

default: install_tools gen-abi gen-proto build

install_tools:
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/bufbuild/buf/cmd/buf@latest

gen-abi:
	@find ./x -name '*.abi.json' | while read abi_file; do \
		pkg_name=$$(basename $$(dirname $$abi_file)); \
		out_file=$$(dirname $$abi_file)/$$(basename $$abi_file .abi.json).abi.go; \
		echo "Generating ABI from $$abi_file into $$out_file"; \
		abigen --abi=$$abi_file --pkg=$$pkg_name --out=$$out_file; \
	done


gen-proto:
	buf lint
	buf format -w
	buf generate

build:
	go mod tidy
	go build cmd/ether-1743613311/main.go