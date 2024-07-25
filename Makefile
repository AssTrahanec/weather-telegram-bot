.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t telegram-bot-weather .

start-container:
	docker run --env-file .env -p 8080:80 telegram-bot-weather
