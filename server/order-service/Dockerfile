FROM golang:1.23.4-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY . .

WORKDIR /app/order-service
RUN go mod download

RUN go build -o main ./cmd/main.go

CMD ["./main"]

