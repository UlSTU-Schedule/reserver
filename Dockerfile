FROM golang:1.16.7-alpine AS builder

WORKDIR /github.com/ulstu-schedule/reserver/

COPY . .

RUN go mod download
RUN go build -o ./.bin/reserver ./cmd/reserver/main.go

FROM alpine:latest

# Install base packages
RUN apk update
RUN apk upgrade
RUN apk add ca-certificates && update-ca-certificates

# Change TimeZone
RUN apk add --update tzdata
ENV TZ=Europe/Samara

# Clean APK cache
RUN rm -rf /var/cache/apk/*

WORKDIR /root/

COPY --from=0 /github.com/ulstu-schedule/reserver/.bin/reserver .
COPY --from=0 /github.com/ulstu-schedule/reserver/configs configs/

CMD ["./reserver"]