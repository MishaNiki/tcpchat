.PHONY: build

build: 
		go build -o serv.exe -v ./cmd/tcpchat

.PHONY: test

test:
		go test -v -race -timeout 30s ./...
			
.DEFAULT_GOAL := build