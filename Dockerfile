# Use the offical golang image to create a binary.
# This is based on Debian and sets the GOPATH to /go.
FROM golang:1.15-alpine AS builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN CGO_ENABLED=0 go build -mod=readonly -v -o server


# Final stage
FROM scratch
COPY --from=builder /app/server /app/server
EXPOSE 8089
ENTRYPOINT ["/app/server"]