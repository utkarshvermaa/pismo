# Base Image
FROM golang:1.22 as builder

# Set working directory
WORKDIR /go/app

# Copy the go.mod and go.sum first
COPY ./go.sum ./go.sum
COPY ./go.mod ./go.mod

# Install Go modules
RUN go mod tidy

# Copy the entire project
COPY . .

# Change user permissions for main.go
RUN chmod +x ./main.go

# Build the app 
RUN GOOS=linux GOARCH=amd64 go build -o txn-routine .

# Runtime image
FROM alpine:3.14

# Set working directory
WORKDIR /root

# Copy the main executable
COPY --from=builder /go/app/txn-routine /app/txn-routine
COPY --from=builder /go/app/migrations /app/migrations


# Set environment variables
ENV SQL_DRIVER=$SQL_DRIVER
ENV SQL_DATA_SRC=$SQL_DATA_SRC
ENV APP_PORT=$APP_PORT
ENV MIGRATION_SRC=$MIGRATION_SRC

# Expose the listening port
EXPOSE 8080

RUN ls -l /app/
RUN ls -l /app/txn-routine
RUN chmod +x /app/txn-routine


# Start the application
ENTRYPOINT ["/app/txn-routine", "start"]
