FROM golang:1.21-alpine

WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod ./
COPY go.sum ./

# Download all Go module dependencies
RUN go mod download

# Copy the entire application source
COPY . .

# Build the Go binary from the cmd directory
RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]
