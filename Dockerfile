FROM golang:1.23 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o main .

FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/main .
COPY wait-for-it.sh .

RUN apt-get update && apt-get install -y \
    ca-certificates \
    netcat-openbsd && \
    chmod +x wait-for-it.sh && \
    rm -rf /var/lib/apt/lists/*

EXPOSE 8082

CMD ["./wait-for-it.sh", "localhost:26257", "--", "./main"]
