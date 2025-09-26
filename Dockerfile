# Build stage
FROM golang:1.21-alpine AS build

WORKDIR /app

# Install git (required for go mod) and build tools
RUN apk add --no-cache git

# Copy go mod and sum, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project (except files in .dockerignore)
COPY . .

# Build the Go app
RUN go build -o siduk cmd/main.go

# Runtime stage
FROM alpine:3.18

WORKDIR /app

# Install CA certificates (for HTTPS) and timezone data if needed
RUN apk add --no-cache ca-certificates tzdata

# Copy binary and static files from build stage
COPY --from=build /app/siduk .
COPY --from=build /app/public ./public
COPY --from=build /app/views ./views

# (Optional) copy config or .env if needed
# COPY .env .env

# Expose the default Fiber port
EXPOSE 8080

# Run the app
CMD ["./siduk"]