FROM golang:1.18-alpine3.16

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/golangtesting main.go

ENTRYPOINT ["/app/golangtesting"]




