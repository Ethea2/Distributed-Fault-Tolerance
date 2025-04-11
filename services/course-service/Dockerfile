FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy common package first
COPY services/common /app/services/common

# Copy this service
COPY services/course-service /app/services/course-service

# Work in the service directory
WORKDIR /app/services/course-service

# Download dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Create a minimal production image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/services/course-service/main .

# Command to run
CMD ["./main"]
