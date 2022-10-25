FROM golang:1.19.2-alpine3.15

WORKDIR /usr/local/incantationChantingServer
COPY src/ /usr/local/incantationChantingServer/src/
COPY go.mod /usr/local/incantationChantingServer/
COPY go.sum /usr/local/incantationChantingServer/
RUN mkdir -p /usr/local/incantationChantingServer/tmp
RUN go mod tidy
RUN go build -o incantationChantingServer /usr/local/incantationChantingServer/src/main.go

CMD ["./incantationChantingServer"]
