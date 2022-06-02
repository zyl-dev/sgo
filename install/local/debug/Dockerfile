FROM golang:alpine AS build-env

ENV CGO_ENABLED 0

# Allow Go to retreive the dependencies for the build step
RUN apk add --no-cache git

WORKDIR /docker-env-debugger/
ADD ./src /docker-env-debugger/

RUN go build -o /docker-env-debugger/srv main.go

# Get Delve from a GOPATH not from a Go Modules project
WORKDIR /go/src/
RUN go get github.com/go-delve/delve/cmd/dlv

# final stage
FROM alpine
WORKDIR /
COPY --from=build-env /docker-env-debugger/srv /
COPY --from=build-env /go/bin/dlv /

EXPOSE 8080 40000

CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/srv"]