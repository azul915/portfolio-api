FROM golang:1.13-alpine

WORKDIR /go/src

COPY . .

ENV GO111MODULE=on

# RUN set -ex && \
#     apk update && \
#     apk add --no-cache git && \
#     # go build -o app &&
#     : "for Gin Web Framework" && \
#     go get github.com/gin-gonic/gin && \
#     : "for Firebase Admin SDK" && \
#     go get firebase.google.com/go