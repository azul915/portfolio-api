FROM golang:1.13-alpine

WORKDIR /go/src/app

COPY . .

RUN set -ex && \
    apk update && \
    apk add --no-cache git && \
    # go build -o app &&
    : "for Firebase Admin SDK" && \
    go get firebase.google.com/go