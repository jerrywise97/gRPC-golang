# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

# Install protoc (cf. http://google.github.io/proto-lens/installing-protoc.html)
ENV PROTOC_ZIP=protoc-3.13.0-linux-x86_64.zip
RUN apt-get update && apt-get install -y unzip
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.13.0/$PROTOC_ZIP \
    && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc \
    && unzip -o $PROTOC_ZIP -d /usr/local 'include/*' \ 
    && rm -f $PROTOC_ZIP

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build https://github.com/jerrywise97/gRPC-golang.git


EXPOSE 9000

CMD ["go", "run", "./server.go"]
