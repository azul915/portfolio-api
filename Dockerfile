FROM golang:1.13.8-alpine as dev

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
CMD ["go", "run", "-v", "server.go"]

FROM golang:1.13.8-alpine as builder

WORKDIR /go/src
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

COPY --from=dev /go/src /go/src
RUN set -ex && \
    go build -a -o /go/bin/api .
CMD ["/go/bin/api"]



FROM alpine:latest as prod

WORKDIR /app
COPY --from=builder /go/bin/api .

CMD [ "/app/api" ]
