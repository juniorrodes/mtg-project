FROM golang:1.22.5-alpine3.19 AS builder

WORKDIR /app

COPY . ./

RUN go build -o mtg-application ./pkg/main.go

FROM alpine:3.19

WORKDIR /mtg-application

COPY --from=builder /app/mtg-application .

CMD ["./mtg-application"]