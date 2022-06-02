# syntax = docker/dockerfile:experimental

FROM golang:alpine

ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
ENV CONFIGOR_ENV develop

ADD ./src /go/src
#RUN apk --no-cache add linux-headers g++ gcc git \
#    && mkdir /go/cache \
#    && mkdir /go/bin/logs /go/bin/configs -p \
#    && cd /go/src/
#    && go mod download \
#    && go build -v -o /go/bin/api ./main.go
WORKDIR /go/src
RUN cd /go/src && go mod download
#FROM alpine as api
#WORKDIR /work
#COPY --from=build /go/bin /work/
#CMD  /work/api
