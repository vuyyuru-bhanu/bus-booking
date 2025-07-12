FROM golang:1.21-alpine

WORKDIR /app

# Copy go.mod (and optionally go.sum if it exists)
COPY go.mod ./

# Download dependencies and generate go.sum
RUN go mod tidy

# Copy the full application code
COPY . .

# Build the Go application
RUN go build -o main ./cmd

EXPOSE 8080

ENTRYPOINT ["./main"]
