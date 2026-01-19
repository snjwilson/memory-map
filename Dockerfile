# Stage 1: Build
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/server 

# Stage 2: Run
FROM alpine:latest
WORKDIR /root/
# Copy the binary from the builder stage
COPY --from=builder /app/main .
# Copy migrations if you have them
COPY --from=builder /app/migrations ./migrations 

EXPOSE 8080
CMD ["./main"]