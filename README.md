# Test Transaction

## Run

`Up postgresql`

``` 
docker-compose -f docker-compose.yml up -d 
```

`Run transaction processor`

```
go run cmd/processor/main.go
```

`Run api server`

```
go run cmd/api/main.go
```

## API

### Transfer Request

`POST /transfer`

    curl -X POST http://localhost:8000/transfer -d '{"user_id": "Eugene", "type": "deposit", "amount": 201230}'

### Response

    HTTP/1.1 200 OK
    Date: Mon, 15 Aug 2022 00:54:13 GMT
    Content-Type: application:json
    Content-Length: 66
    
    "6eff69a580efcb67f723b5d3e508ae760b0be4fef911002c012f0cd54aabbc1e"

### Transaction Info

`Get /transaction/{transaction_key}`

    curl -i http://localhost:8000/transaction/d87d81efdfa437a02408fce3e4884650bbf1c625c0d2cd649bb70730a81540e4

### Response

    HTTP/1.1 200 OK
    Date: Mon, 15 Aug 2022 00:57:27 GMT
    Content-Type: application:json
    Content-Length: 105
    
    {"UserID":"Eugene","Amount":201230,"CreatedAt":"2022-08-15T03:53:02.389688Z","Status":"pending"}
    
### TODO

- [ ] Dockerfile's for both services
- [ ] Api to get user balance
