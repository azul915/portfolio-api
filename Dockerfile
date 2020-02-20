FROM golang:latest

WORKDIR /go/src

ENV GO111MODULE=on

ADD ./api ./api
ADD server.go .
ADD credentials.json .
ADD go.mod .
ADD go.sum .
