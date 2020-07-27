FROM golang:1.14.6-alpine

ARG app_env
ENV APP_ENV $app_env
ENV CGO_ENABLED 0

RUN apk update && apk add bash inotify-tools git coreutils

RUN go get github.com/go-delve/delve/cmd/dlv

COPY . /go/src/github.com/lilahamstern/ced/server
WORKDIR /go/src/github.com/lilahamstern/ced/server

RUN chmod +x startScript.sh

ENTRYPOINT sh startScript.sh

EXPOSE 8000 2345