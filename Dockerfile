FROM golang:alpine

MAINTAINER Maintainer

ENV GIN_MODE=release
ENV PORT=3004

WORKDIR /go/src/rfc5322fun

COPY . /go/src/rfc5322fun/
RUN go build -o ./app /go/src/rfc5322fun/

EXPOSE $PORT

ENTRYPOINT ["./app"]
