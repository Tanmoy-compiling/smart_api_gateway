# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the product service
RUN CGO_ENABLED=0 GOOS=linux go build -o /product-service ./services/product

# Stage 2: Create the final image
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /product-service .

# Expose the port the application runs on
EXPOSE 5002

# Command to run the application
CMD ["./product-service"]