FROM golang:1.20 AS builder

WORKDIR /app

ADD . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest as production

COPY --from=builder /app .

CMD [ "./app" ]