tools-generate:
	go install github.com/matryer/moq@latest

tools-lint: tools-generate
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

generate:
	go generate ./...

test: generate
	go test -cover -v ./...

lint: generate
	golangci-lint run

run:
	docker-compose up -d --build app

build:
	cd cmd/yaus && go build .
