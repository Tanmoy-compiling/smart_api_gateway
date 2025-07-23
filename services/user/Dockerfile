# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the user service
RUN CGO_ENABLED=0 GOOS=linux go build -o /user-service ./services/user

# Stage 2: Create the final image
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /user-service .

# Expose the port the application runs on
EXPOSE 5001

# Command to run the application
CMD ["./user-service"]