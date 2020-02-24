FROM golang:1.13.8-alpine as builder

WORKDIR /go/src
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

COPY go.mod .
COPY go.sum .
RUN set -ex && \
    apk add --no-cache git && \
    go mod download

COPY ./api ./api
COPY server.go .
COPY credentials.json .
RUN set -ex && \
    go build -a -o /go/bin/api .




FROM alpine:latest as prod

WORKDIR /app
COPY --from=builder /go/bin/api .
CMD ["/app/api"]




FROM golang:1.13.8-alpine as dev

WORKDIR /go/src
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
COPY ./api ./api
COPY server.go .
COPY credentials.json .
COPY .realize.yaml .

RUN set -ex && \
    apk update && \
    apk add --no-cache git && \
    : "for Gin Web Framework" && \
    go get github.com/gin-gonic/gin && \
    : "for Firebase Admin SDK" && \
    go get firebase.google.com/go && \
    : "for realize hot reloader" && \
    go get github.com/oxequa/realize

EXPOSE 1999
CMD [ "realize", "start" ]
