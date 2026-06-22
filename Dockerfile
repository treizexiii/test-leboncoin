FROM golang:latest AS BUILD

COPY . /app
WORKDIR /app

RUN go mod tidy
RUN go build -o dist/main .
COPY config.yaml dist/config.yaml

FROM debian:latest

WORKDIR /app
COPY --from=BUILD /app/dist .

CMD ["./main"]
