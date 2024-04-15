FROM golang:latest

RUN go mod download
RUN go build 

# your port of choice, make sure it matches with the port in your .env 
EXPOSE 13234

ENTRYPOINT ["go", "run", "server.go"]
