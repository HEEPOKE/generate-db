# generate-db

## Config Environment

```bash
cp .env.example .env
```

## Step Build Docker && Run

```bash
docker network create Heepoke
```

```bash
docker compose up -d && docker compose logs api --follow
```

## Create DB For test

```bash
docker compose -f db-test.yml up -d && docker compose -f db-test.yml logs
```

## Test

```bash
go test ./test/... -v
```

- coverage test

```bash
go test -cover ./test/... -v
```

## Generate Swagger

```bash
swag init -g cmd/main.go --output=pkg/docs
```

## Local Swagger Doc Api

<http://localhost:6476/apis/docs/index.html>

## Example DSN Database

- mysql

```bash
usertest:test@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=true&loc=Local
```

- postgresql

```bash
host=db-postgres-test user=test dbname=test password=test sslmode=disable TimeZone=Asia/Bangkok
```

- sqlserver

```bash
sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm
```

- mongodb

```bash
mongodb://test:test@localhost:27017/test
```
