mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/unbox && \
	go mod tidy

run:
	cd cmd/unbox/ && \
	go run .

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

check: fmt lint

go-build:
	cd cmd/unbox/ && \
	CGO_ENABLED=0 go build -o ../../tmp/unbox .

brun: go-build
	cd tmp && \
	./unbox