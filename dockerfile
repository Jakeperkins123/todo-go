# Use a proper Go version
FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/dist .

# Use a smaller base image for the runtime
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /app/out/dist /app/dist

# Set the working directory and command
WORKDIR /app
CMD ["./dist"]
