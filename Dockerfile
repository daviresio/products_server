# Build Stage
FROM golang:alpine3.20 AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/app

# Cache dependencies based on go.{mod,sum} files
COPY go.mod go.sum ./

RUN go mod download

# Copy the remaining application files
COPY ./ ./

# Run unit tests
RUN CGO_ENABLED=0 go test -v ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/app ./cmd

# Production Image Stage
FROM alpine:3.20.3

RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=build_base /tmp/app/out/app /app

# Expose the service port
EXPOSE 8080

# Run the binary program
CMD ["/app/app"]
