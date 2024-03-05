# Use the official Golang image as the base image
FROM golang:1.17 as builder

# Set the working directory to /app
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Install the dependencies
RUN go mod download

# Copy the source code to the working directory
COPY *.go ./

# Build the Go binary
RUN go build -o main .

# Create a new stage for running the app
FROM alpine:3.14

# Set the working directory to /app
WORKDIR /app

# Copy the main binary to the working directory
COPY --from=builder /app/main /app

# Set the default command to run the app
CMD ["/app/main"]