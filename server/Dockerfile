FROM golang:1.20.2-alpine3.17

WORKDIR /go/src/server
COPY . /go/src/server

RUN apk update && \
    apk add --no-cache git gcc musl-dev make

RUN go install golang.org/x/tools/gopls@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 9090
