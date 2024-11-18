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
	@find ./contracts -name '*.abi' | while read abi_file; do \
		pkg_name=$$(basename $$(dirname $$abi_file)); \
		out_file=$$(dirname $$abi_file)/$$(basename $$abi_file .abi).abi.go; \
		echo "Generating ABI for $$abi_file into $$out_file"; \
		abigen --abi=$$abi_file --pkg=$$pkg_name --out=$$out_file; \
	done