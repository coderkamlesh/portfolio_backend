# Stage 1: Build stage
# Humne 1.25-alpine use kiya hai taaki tere go.mod ke version se match kare
FROM golang:1.25-alpine AS builder

# Git aur build essentials install karein (kuch Go modules ke liye zaroori hote hain)
RUN apk add --no-cache git

WORKDIR /app

# Pehle dependencies copy karo (Caching benefit)
COPY go.mod go.sum ./
RUN go mod download

# Ab poora code copy karo
COPY . .

# Binary build karo. Path cmd/api/main.go hai tere project mein
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# Stage 2: Final Light Image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Builder se sirf binary aur .env file uthao
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Gin default port 8080 use karta hai
EXPOSE 8080

# Run the binary
CMD ["./main"]