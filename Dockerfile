FROM golang:1.23.4-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o telegram-bot-crm .
FROM alpine:latest
RUN apk add --no-cache sqlite
COPY --from=builder /app/telegram-bot-crm /usr/local/bin/
EXPOSE 8080
CMD ["telegram-bot-crm"]
