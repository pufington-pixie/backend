# Use the official Go image as the base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod download

# Copy the project files to the working directory
COPY . .

# Build the Go application
RUN go build -o main ./cmd

# Expose the port on which your application will run
EXPOSE 8080

# Run the Go application
CMD ["./main"]
