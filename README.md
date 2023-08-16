# generate-db

## Config Environment

```bash
cp .env.example .env
```

## Step Build Docker && Run

```bash
docker create network Heepoke
```

```bash
docker compose up -d && docker compose logs api --follow
```

## Create DB For test

```bash
docker compose -f db-test.yml up -d && docker compose -f db-test.yml logs --follow
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

```bash
http://localhost:6476/apis/docs/index.html
```

## Example DSN Database

- mysql

```bash
user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
```

- postgresql

```bash
host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
```

- sqlserver

```bash
sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm
```

- mongodb

```bash
mongodb://test:test@localhost:27017/mongo-db-test
```
