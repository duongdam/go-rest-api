# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.18 as builder

ADD ./ /go/src/app
# Create and change to the app directory.
WORKDIR /go/src/app
# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
# Copy local code to the container image.
# export some environment file in linux then use it in project, example environment
COPY .env.production .env

RUN go get ./
RUN go build -o server
# Build the binary.

# Run the web service on container startup.
CMD ["./server"]
