FROM golang:1.22.2

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go get

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# your port of choice, make sure it matches with the port in your .env 
EXPOSE 13234

CMD ["/docker-gs-ping"]
