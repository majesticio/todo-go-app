# Use the official Golang image as a builder
FROM golang:1.21 AS builder

# Set the working directory
WORKDIR /src

# Copy local code to the container image.
COPY . .

# Build a static binary.
RUN CGO_ENABLED=0 go build -o /bin/app -a -tags netgo -ldflags '-w -extldflags "-static"'

# Use a minimal Alpine image for the final stage
FROM alpine:3.15

# Copy the binary to the production image from the builder stage.
COPY --from=builder /bin/app /bin/app

# Run the web service on container startup.
CMD ["/bin/app"]
