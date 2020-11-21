FROM golang:1.15-alpine as builder

WORKDIR /

RUN apk add --no-cache git

ENV GOBIN=/go/bin

ENV GOPATH=

ENV GOROOT=/usr/local/go

COPY ./go.mod ./go.sum ./main.go ./

RUN go mod download

RUN env GOOS=linux GOARCH=amd64 GIN_MODE=release go build -o /go-api

FROM alpine:3.12.0

WORKDIR /myapp

COPY --from=builder /go-api .

CMD [ "./go-api" ]