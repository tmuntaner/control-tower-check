.PHONY: all
all: tidy build

.PHONY: tidy
tidy:
	go mod tidy
	go fmt ./cmd/... ./internal/...

.PHONY: build
build:
	go build -o ./control-tower-check ./cmd/main.go
