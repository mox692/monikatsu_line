FROM golang:1.15-alpine as builder

WORKDIR /go/src/github.com/mox692/manage_sesstion/session/rest/server

RUN apk add --no-cache git

ENV GOBIN=/go/bin

ENV GOPATH=/go

ENV GO111MODULE=on

ENV GOROOT=/usr/local/go

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./ ./

RUN env GOOS=linux GOARCH=amd64 go build -o /session ./server/server.go 
# RUN chmod 777 ./server/server.go

FROM alpine:3.12.0

WORKDIR /myapp

# envfileもalpineコンテナに入れておく。
COPY --from=builder /session ./ 

CMD ["./session"]