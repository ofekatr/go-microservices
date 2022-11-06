FROM golang:1.18-alpine as builder 

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o broker_app ./cmd/api

RUN chmod +x broker_app


FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/broker_app /app

CMD ["/app/broker_app"]