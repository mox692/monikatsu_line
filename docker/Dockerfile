FROM golang:1.15-alpine as builder

WORKDIR /go/src/github.com/mox692/monikatsu_line

RUN apk add --no-cache git

ENV GOBIN=/go/bin

ENV GOPATH=/go

ENV GO111MODULE=on

ENV GOROOT=/usr/local/go


COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./ ./

RUN env GOOS=linux GOARCH=amd64 go build -o /go-api

FROM alpine:3.12.0

WORKDIR /myapp

EXPOSE 8080

# envfileもalpineコンテナに入れておく。
COPY --from=builder /go-api /go/src/github.com/mox692/monikatsu_line/local.env ./ 

CMD [ "./go-api" ]