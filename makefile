all: .build

.PHONY: build

.build: .format
	go mod tidy

.format:
	go mod tidy
	gofmt -w .
	goimports -w .
	golint ./...
	go-xray -d .
	gonote .