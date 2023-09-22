# Use the official Go image as a base
FROM golang:1.21.1

# Set up the working directory
WORKDIR /go/src/app

# Enable Go modules
ENV GO111MODULE=on

# Copy go mod and sum files to download dependencies
COPY go.mod go.sum ./

# Download all dependencies. 
# They will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the entire source code from the current directory to the container
COPY . .

# Expose port 3000 to the outside
EXPOSE 3000

# Command to run the application
CMD ["go", "run", "main.go"]
