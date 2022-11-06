FROM golang:1.18-alpine as builder 

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o front_app ./cmd/web

RUN chmod +x front_app


FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/front_app .

RUN mkdir -p ./cmd/web/templates

COPY --from=builder /app/cmd/web/templates ./cmd/web/templates

CMD ["./front_app"]