# generate-db

## Config Environment

```bash
 cp .env.example .env
```

## Build Docker

```bash
docker compose up -d
```

## How to run

```bash
go run cmd/main.go
```

## Generate Swagger

```bash
swag init -g cmd/main.go --output=pkg/docs
```

## Local Swagger Doc Api

```bash
http://localhost:6476/apis/docs/index.html
```
