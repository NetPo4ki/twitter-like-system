# Use the official Golang image for the build stage
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy all files from the current directory to the container
COPY . .

# Download and cache Go modules
RUN go mod download

# Build the application
RUN go build -o user-service cmd/main.go

CMD ["/app/user-service"]