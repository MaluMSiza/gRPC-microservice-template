# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /grpc-server ./cmd/server/main.go
RUN go build -o /grpc-client ./cmd/client/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /grpc-server /grpc-client ./
COPY --from=builder /app/proto ./proto

EXPOSE 50051

CMD ["./grpc-server"]