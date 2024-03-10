FROM golang:1.21.1-alpine3.18 AS builder

RUN apk update && \
    apk add bash ca-certificates git gcc g++ libc-dev binutils file

WORKDIR /opt

# COPY go.mod go.sum ./
COPY . .
RUN go mod download && go mod verify
RUN go build -o /opt/wcharge_mqtt ./cmd/wcharge_mqtt/main.go


FROM alpine:3.18 AS production
RUN apk update && \
    apk add ca-certificates libc6-compat && rm -rf /var/cache/apk/*

WORKDIR /opt

COPY --from=builder /opt/wcharge_mqtt ./
COPY --from=builder /opt/config/* ./config/


CMD ["./wcharge_mqtt"]
