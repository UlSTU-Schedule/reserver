.PHONY:
.SILENT:

build:
	go build -o ./.bin/reserver cmd/reserver/main.go

run: build
	./.bin/reserver

build-image:
	docker build -t ulstu-schedule/reserver .

start-container:
	docker run --env-file .env --rm ulstu-schedule/reserver
