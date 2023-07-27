FROM golang:1.20-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .


EXPOSE 6476

CMD ["air"]