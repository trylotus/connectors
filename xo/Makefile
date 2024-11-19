gen-abi:
	abigen --abi=./xo/xo.abi.json --pkg=xo --out=./xo/xo.abi.go

gen-proto:
	buf lint
	buf format -w
	buf generate
