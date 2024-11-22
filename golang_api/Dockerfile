# Use Go 1.23.3 base image
FROM golang:1.23.3

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy application code
COPY . .

# Expose ports (for golang_api only)
EXPOSE 8080

# Default command (to be overridden by docker-compose.yml if needed)
CMD ["go", "run", "main.go"]
