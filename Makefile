mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/unbox && \
	go mod tidy

run:
	go run . \
		-f tmp/VLESS.txt \
 		-t tmp/tmpl.json \
 		-o tmp/sing-box.json -j yes

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

check: fmt lint