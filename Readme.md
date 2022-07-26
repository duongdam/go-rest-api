## Start API

```bash
go get .
go run .
```

## Test call api

```bash
curl --location --request GET 'http://localhost:8080/api/user/list'

// output
[
    {
        "id": "1-duong",
        "name": "Duong",
        "address": "Nga Thuy - Nga Son - Thanh Hoa"
    },
    {
        "id": "2-phi",
        "name": "Phi",
        "address": "Nga Trung - Nga Son - Thanh Hoa"
    },
    {
        "id": "3-thanh",
        "name": "Thanh",
        "address": "TT Nga Son - Thanh Hoa"
    }
]

curl --location --request GET 'http://localhost:8080/api/user/get/1-duong' \
--data-raw ''

// output

{
    "id": "1-duong",
    "name": "Duong",
    "address": "Nga Thuy - Nga Son - Thanh Hoa"
}


```