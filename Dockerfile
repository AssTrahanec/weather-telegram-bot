FROM golang:1.22.2-alpine AS builder


COPY . /github.com/asstrahanec/weather-telegram-bot/
WORKDIR /github.com/asstrahanec/weather-telegram-bot/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/asstrahanec/weather-telegram-bot/.bin/bot .
COPY --from=0 /github.com/asstrahanec/weather-telegram-bot/configs configs/

EXPOSE 80

CMD ["./bot"]