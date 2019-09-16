## Rule sets in stub repository

### Organisation 1
* Is greater than 5 and less than 5000

### Organisation 2
* Is greater than 500 and less than 1000

## Getting started

```bash
$ go build main.go && ./main -> listening on port 8080
```

POST localhost:8080/validate

### Example 1

request:
```json
{
	"amount": 4000,
	"organization": "1"
}
```

response:
```json
{
    "action": "PASS",
    "block": [],
    "tags": []
}
```

### Example 2

request:
```json
{
	"amount": 900,
	"organization": "1"
}
```

response:
```json
{
    "action": "BLOCK",
    "block": [
        {
            "name": "Is greater than 5 and less than 5000",
            "rules": [
                {
                    "key": "amount",
                    "operator": "<",
                    "value": "5000"
                },
                {
                    "key": "amount",
                    "operator": ">",
                    "value": "5"
                }
            ]
        }
    ],
    "tags": []
}
```

### Example 3

request:
```json
{
	"amount": 900,
	"organization": "2"
}
```

response:
```json
{
    "action": "PASS",
    "block": [],
    "tags": [
        {
            "name": "Is greater than 500 and less than 1000",
            "rules": [
                {
                    "key": "amount",
                    "operator": "<",
                    "value": "1000"
                },
                {
                    "key": "amount",
                    "operator": ">",
                    "value": "500"
                }
            ]
        }
    ]
}
```
