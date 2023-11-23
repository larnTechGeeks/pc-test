FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN chmod -R 755 /app/templates

RUN go build -o spamService cmd/api/*

FROM alpine:latest

WORKDIR /usr/bin

COPY --from=builder /app/spamService .
COPY --from=builder /app/docs .
COPY --from=builder /app/templates /templates

CMD [ "spamService" ]