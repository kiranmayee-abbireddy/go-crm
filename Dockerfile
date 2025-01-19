# Use the official Golang image to build the application
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Install dependencies
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o crm-app .

# Final stage: use a minimal Alpine image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built app from the previous stage
COPY --from=builder /app/crm-app .

# Expose the port on which the app will run
EXPOSE 8081

# Run the Go app
CMD ["./crm-app"]
