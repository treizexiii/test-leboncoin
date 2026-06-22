FROM golang:latest AS BUILD

COPY . /app
WORKDIR /app

RUN go mod tidy
RUN go build -o dist/main . 

FROM debian:latest

WORKDIR /app
COPY --from=BUILD /app/dist/main .
EXPOSE 8080

CMD ["./main"]
