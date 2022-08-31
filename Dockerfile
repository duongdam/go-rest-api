# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.19-alpine as build-stage

# Create and change to the app directory.
WORKDIR /dist
# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
# Copy local code to the container image.
# export some environment file in linux then use it in project, example environment
COPY . .
COPY .env.production .env
RUN go get ./
RUN go build -o server

# Stage 2: build the image
FROM alpine
WORKDIR /app
COPY --from=build-stage /dist/.env ./.env
COPY --from=build-stage /dist/server ./server
COPY --from=build-stage /dist/static ./static
CMD ["./server"]
