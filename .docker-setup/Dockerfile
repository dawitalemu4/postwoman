FROM dawitalemu4/postwoman:latest AS builder


FROM golang:1.22.2

RUN apt-get update && apt-get install -y curl

COPY --from=builder /postwoman /postwoman
COPY --from=builder /go/views /go/views

COPY .env .

CMD ["/postwoman"]
