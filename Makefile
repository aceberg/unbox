mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/unbox && \
	go mod tidy

run:
	go run . \
		-f tmp/VLESS.txt \
 		-t tmp/tmpl.json \
 		-o tmp/sing-box.json -j

arun:
	go run . \
		-a "http://127.0.0.1:9091" \
 		-o tmp/sing-box.json

drun:
	go run . \
 		-o tmp/sing-box.json -d

krun:
	go run . \
		-a "http://127.0.0.1:9091" \
 		-k

irun:
	go run . \
 		-i tmp/sing-box.json

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

check: fmt lint