FROM golang:1.21-alpine

WORKDIR /app

# Copy go.mod only (no go.sum yet)
COPY go.mod ./

# Initialize an empty go.sum by forcing Go to resolve dependencies
COPY . .  # Copy the full source now (since go.sum doesn't exist yet)

# Generate go.sum and download modules
RUN go mod tidy

# Build the app
RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]
