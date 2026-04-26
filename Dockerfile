FROM golang:alpine AS builder

RUN apk add build-base
COPY . /src
RUN cd /src/cmd/unbox/ && CGO_ENABLED=0 go build -o /unbox .


FROM scratch

WORKDIR /data/unbox
WORKDIR /app

COPY --from=builder /unbox /app/

ENTRYPOINT ["./unbox"]