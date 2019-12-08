FROM golang:1.13-alpine as build

WORKDIR /go/app

COPY . .

RUN set -ex && \
    apk update && \
    apk add --no-cache git && \
    # go build -o app && 
    : "for hot reload" && \
    go get gopkg.in/urfave/cli.v2@master && \
    go get github.com/oxequa/realize && \
    : "for Firebase Admin SDK" && \
    go get firebase.google.com/go

FROM alpine

WORKDIR /app

COPY --from=build /go/app/app .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /app/app

CMD ["./app"]