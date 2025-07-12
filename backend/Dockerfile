FROM golang:1.21-alpine

WORKDIR /app

# Copy everything (go.mod, .go files, etc.)
COPY . .

# Generate go.sum and download all dependencies
RUN go mod tidy

# Build the app
RUN go build -o main ./cmd

EXPOSE 8080

ENTRYPOINT ["./main"]
