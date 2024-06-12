all: .build

.PHONY: build

.build: .format
	go mod tidy
	protoc --proto_path=${GOPATH}/src/:${GOPATH}/src/github.com/trpcprotocol:. --go-grpc_out=require_unimplemented_servers=false:. --trpc2grpc_out=require_unimplemented_servers=false:. --go_out=. --grpc-gateway_out=. cache_server.proto
	mv -f github.com/abetterchoice/protoc_cache_server/* ./
	rm -rf github.com

.format:
	go mod tidy
	gofmt -w .
	goimports -w .
	golint ./...
	gonote ./...