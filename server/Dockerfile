# Use the official Golang image as the base image
FROM golang:1.23 as builder

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/api
RUN ls -l /app

# Create a smaller image for the final production stage
FROM alpine:3.17

# Install the golang-migrate tool
RUN apk add --no-cache curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate /usr/local/bin/migrate

WORKDIR /app

# Copy the server binary from the builder stage
COPY --from=builder /app/server /app/server

COPY entrypoint.sh /app/entrypoint.sh

# Copy any migration files
COPY ./cmd/migrate/migrations /app/migrations

RUN chmod +x /app/entrypoint.sh

CMD ["/bin/sh", "-c", "/app/entrypoint.sh && ./server"]
