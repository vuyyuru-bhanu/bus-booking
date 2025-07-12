FROM golang:1.21-alpine

WORKDIR /app

# Copy only go.mod first
COPY go.mod ./

# Generate go.sum (even if it doesn’t exist yet)
RUN go mod tidy

# Now copy the full source code
COPY . .

# Build the Go app
RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]
