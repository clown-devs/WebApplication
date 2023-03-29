.PHONY: build, run, all, check

build: 
	go build  -o ./bin/server -v ./cmd/main.go 
	
run:
	go run ./cmd/main.go

check:
	go vet ./...

	
.DEFAULT_GOAL := run