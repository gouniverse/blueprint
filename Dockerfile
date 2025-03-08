# 1. Use the official Golang image to create a build artifact.
# ===============================================================
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.24.0 as builder

# Create and change to the app directory.
WORKDIR /app

# 2. Retrieve application dependencies using go modules.
# ===============================================================
# Allows container builds to reuse downloaded dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# 3. Build the binary.
# ===============================================================
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
# RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server

RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

# 4. Use the official Alpine image for a lean production container.
# ===============================================================
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ="UTC"

# Alternative using Debian Slim
# =============================================================== 
# FROM debian:buster-slim
# RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
#     ca-certificates && \
#     rm -rf /var/lib/apt/lists/*

# 5. Copy the binary to the production image from the builder stage.
# ===============================================================
COPY --from=builder /app/server /server

# COPY --from=builder /app/views /views

# Run the web service on container startup.
CMD ["/server"]
