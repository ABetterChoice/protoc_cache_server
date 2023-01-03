all: .build

.PHONY: build

.build: .format
	go mod tidy
	protoc --proto_path=$GOPATH/src:${GOPATH}/src/git.code.oa.com/trpcprotocol:. --go-grpc_out=require_unimplemented_servers=false:. --trpc2grpc_out=require_unimplemented_servers=false:. --go_out=. --grpc-gateway_out=. cache_server.proto
	mv -f git.woa.com/tencent_abtest/protocol/protoc_cache_server/* ./
	rm -rf git.woa.com

.format:
	go mod tidy
	gofmt -w .
	goimports -w .
	golint ./...
	go-xray -d .
	gonote .