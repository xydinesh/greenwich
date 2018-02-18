FROM golang:alpine
MAINTAINER Dinesh Weerapurage <dweerapurage@mapanything.com>
RUN apk update && apk add git make ca-certificates unzip curl && update-ca-certificates
# install dep
WORKDIR /root
RUN curl -L -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64
WORKDIR /go/src
RUN git clone https://github.com/xydinesh/greenwich.git
WORKDIR greenwich
RUN git checkout dev
RUN make
