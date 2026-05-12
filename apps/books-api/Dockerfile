# 1. Build Stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy modules
COPY go.mod ./
# Copy source
COPY main.go ./

# Compile a static Go binary (CGO_ENABLED=0 ensures it runs anywhere)
RUN CGO_ENABLED=0 GOOS=linux go build -o books-api main.go

# 2. Run Stage
# Using Google's "distroless" image for ultimate security and tiny footprint
FROM gcr.io/distroless/static-debian11

WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /app/books-api .

EXPOSE 8080

# Run as a non-root user for security best practices
USER nonroot:nonroot

ENTRYPOINT ["/books-api"]
