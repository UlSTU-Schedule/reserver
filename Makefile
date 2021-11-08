.PHONY:
.SILENT:

build:
	go build -o ./.bin/reserver main.go

run: build
	./.bin/reserver

build-image:
	docker build -t tmrrwnxtsn/ulstu-schedule-reserver .

start-container:
	docker run --env-file .env --rm tmrrwnxtsn/ulstu-schedule-resever
