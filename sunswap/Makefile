gen-abi:
	abigen --abi=./erc20/erc20.abi --pkg=erc20 --out=./erc20/erc20.abi.go
	abigen --abi=./factory/factory.abi --pkg=factory --out=./factory/factory.abi.go
	abigen --abi=./pair/pair.abi --pkg=pair --out=./pair/pair.abi.go

gen-proto:
	buf lint
	buf format -w
	buf generate
