FROM golang:1.20.2-alpine3.17

WORKDIR /go/src/pb
COPY . /go/src/pb/

RUN apk update && \
    apk add --no-cache git gcc musl-dev make

RUN go install github.com/bufbuild/buf/cmd/buf@latest && \
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
