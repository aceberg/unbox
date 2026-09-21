FROM --platform=$BUILDPLATFORM tonistiigi/xx AS xx

FROM --platform=$BUILDPLATFORM golang:alpine AS builder

COPY --from=xx / /

WORKDIR /src

RUN apk add --no-cache ca-certificates build-base

COPY . /src
RUN go mod download

ARG TARGETPLATFORM
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /unbox .


FROM scratch

WORKDIR /data/unbox
WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /unbox /app/

ENTRYPOINT ["./unbox"]