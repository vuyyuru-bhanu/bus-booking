FROM golang:1.21-alpine

WORKDIR /app

# Copy go.mod and go.sum (if it exists)
COPY go.mod ./
COPY go.sum* ./

# Run go mod tidy to ensure all dependencies are resolved and go.sum is generated
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]