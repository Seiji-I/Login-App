FROM golang:1.16.9-alpine3.14

ENV GOPATH /go
ENV GO111MODULE on

RUN apk update && \
    apk --no-cache add git

WORKDIR /go/src

ADD . /go/src

RUN go mod tidy && \
    go install github.com/cosmtrek/air@v1.27.3

CMD ["air", "-c", ".air.toml"]