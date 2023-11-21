FROM golang:1.20-buster as task-builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v ./cmd/main.go

FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=task-builder /app/main /app/main
COPY --from=task-builder /app/config /app/config

CMD ["/app/main"]
