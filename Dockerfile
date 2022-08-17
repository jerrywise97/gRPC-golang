# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build https://github.com/jerrywise97/gRPC-golang.git


EXPOSE 9000

CMD ["go", "run", "./server.go"]
