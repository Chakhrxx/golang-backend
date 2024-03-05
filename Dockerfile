# Use a base image with Go installed
FROM golang:latest as develop

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY ./app .
COPY go.mod go.sum /app/

# Build the Go application
RUN go build -o main .

# Expose the port your Go application listens on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
