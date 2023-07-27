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

## Generate Swagger

```bash
swag init -g cmd/main.go --output=pkg/docs
```

## Local Swagger Doc Api

```bash
http://localhost:6476/apis/docs/index.html
```
