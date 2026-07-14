mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/unbox && \
	go mod tidy

run:
	go run . parse \
		-f tmp/VLESS.txt \
 		-t tmp/tmpl.json \
 		-o tmp/sing-box.json -j

arun:
	go run . conf \
		-a "http://127.0.0.1:9091" \
 		-o tmp/sing-box.json

drun:
	go run . conf \
 		-o tmp/sing-box.json -d

krun:
	go run . keep \
		-a "http://127.0.0.1:9091"

irun:
	go run . conf \
 		-i tmp/sing-box.json

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

check: fmt lint