# for you

# FROM dawitalemu4/postwoman:latest AS builder


# FROM golang:1.22.2

# RUN apt-get update && apt-get install -y curl

# COPY --from=builder /postwoman /postwoman
# COPY --from=builder /go/views /go/views

# COPY .env .

# CMD ["/postwoman"]


# for me (push to docker hub)

# FROM golang:1.22.2 AS builder

# COPY . .

# RUN go build -o /postwoman

# docker image build -t postwoman .
# docker image tag postwoman dawitalemu4/postwoman:latest
# docker push dawitalemu4/postwoman:latest


# for me (test locally)

FROM golang:1.22.2 AS builder

COPY . .

RUN go build -o /postwoman


FROM golang:1.22.2

RUN apt-get update && apt-get install -y curl

COPY --from=builder /postwoman /postwoman
COPY --from=builder /go/views /go/views

COPY .env .

CMD ["/postwoman"]
