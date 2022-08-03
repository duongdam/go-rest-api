## Start API

```bash
# Install
go get .

# Run
go run .

or

go build -o server && ./server
```

## Test call api

```bash
curl --location --request GET 'http://localhost:8080/api/user/list'

// output
[
    {
        "id": "1-kenzo",
        "name": "Kenzo",
        "address": "Tokyo - Japan"
    },
    {
        "id": "2-hinata",
        "name": "Hinata",
        "address": "Tokyo - Japan"
    },
    {
        "id": "3-yumi",
        "name": "Yumi",
        "address": "Tokyo - Japan"
    }
]

curl --location --request GET 'http://localhost:8080/api/user/get/2-hinata' \
--data-raw ''

// output

{
    "id": "2-hinata",
    "name": "Hinata",
    "address": "Tokyo - Japan"
}


```

## Test build on docker

- Install docker: https://docs.docker.com/get-docker/

- Build docker images command: ```docker build -t api-brother .```

- Run docker images command: ```docker run api-brother```

- List images : ```docker images```
