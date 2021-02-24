# Validation Service

### Getting Started

```bash
$ git clone git@github.com:dimebox/validation-service.git
```

### Rule sets in stub repository

#### Entity 1

- Is greater than 5 and less than 5000

#### Entity 2

- Is greater than 500 and less than 1000

### Run in development

#### With Go

```bash
$ go build main.go && ./main -> listening on port 8080
```

POST localhost:8080/transaction/validate

#### With Docker

```bash
$ docker-compose up
```

With Docker you don't need .env file. Configurations are within.

POST localhost:3000/transaction/validate

#### Example 1

request:

```json
{
  "amount": 4000,
  "entity": "1"
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

#### Example 2

request:

```json
{
  "amount": 900,
  "entity": "1"
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

#### Example 3

request:

```json
{
  "amount": 900,
  "entity": "2"
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

### Deployment

#### Build image

```bash
$ docker build ./
```

This will build the API image. It will connect to a MondoDB that should run on the same host network (not Docker network).

#### Run container

##### Using environment file:

```bash
$ docker run -p [SERVER PORT]:8080 \
    --mount type=bind,source=/var/log/dimebox,target=/logs \
    --env-file=[ENV FILE NAME] \
    [IMAGE ID]
```

or:

##### Passing environment vars in run command:

```bash
$ docker run -p [SERVER PORT]:8080 \
    --mount type=bind,source=/var/log/dimebox,target=/logs \
    -e MONGO_URL=[VALUE] \
    -e MONGO_DB=[VALUE] \
    -e MONGO_DB_RETRYMILLISECONDS=[VALUE] \
    -e LOG_FILE_MAX_SIZE=[VALUE] \
    -e LOG_ROTATING_COUNT=[VALUE] \
    -e LOG_ROTATING_PERIOD=[VALUE] \
    [IMAGE ID]
```

Note: The path `/var/log/dimebox` must exist in host. Edit it if you need.

##### Examples of environment variables:

```text
MONGO_URL=mongodb://host.docker.internal:27017
MONGO_DB=validationService
MONGO_DB_RETRYMILLISECONDS=0
LOG_FILE_MAX_SIZE=600mb
LOG_ROTATING_COUNT=30
LOG_ROTATING_PERIOD=1d
```
