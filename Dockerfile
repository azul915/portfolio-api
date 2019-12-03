FROM golang:1.13-alpine as build

WORKDIR /go/app

COPY . .

RUN set -ex && \
    apk update && \
    apk add --no-cache git && \ 
    go build -o app && \
    go get gopkg.in/urfave/cli.v2@master && \
    : "for hot reload" && \
    go get github.com/oxequa/realize && \
    : "for debugger" && \
    go get github.com/go-delve/delve/cmd/dlv && \
    go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

FROM alpine

WORKDIR /app

COPY --from=build /go/app/app .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /app/app

CMD ["./app"]