FROM golang:1.12-alpine3.9

MAINTAINER fujiwaram

ENV APP github.com/fujiwaram/gopher-translator
ENV APP_ROOT $GOPATH/src/$APP

ADD . $APP_ROOT
WORKDIR $APP_ROOT

RUN go install $APP
