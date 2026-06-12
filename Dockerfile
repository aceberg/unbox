FROM golang:alpine AS builder

RUN apk add build-base
COPY . /src
RUN cd /src && CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /unbox .


FROM scratch

WORKDIR /data/unbox
WORKDIR /app

COPY --from=builder /unbox /app/

ENTRYPOINT ["./unbox"]