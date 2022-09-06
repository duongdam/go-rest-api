## Start API

```bash
Fast refresh golang app in development

install air: go install github.com/cosmtrek/air@latest

then Reload terminal

run command: air init

then run command to start server: air

Development evironment auto refresh when update code
```

```bash
# Install
go get .

# Run
go run .

or

go build -o server && ./server
```

## Test build on docker

```html
- Install docker: https://docs.docker.com/get-docker/

- Build docker images command: docker build -t api-test . or docker build -t api-test -f Dockerfile-stg .

- Run docker images command: docker run api-test or docker run --publish 8080:8080 api-test (truy cập localhost:8080 để test)

- List images : docker images

```
