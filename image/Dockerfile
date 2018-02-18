FROM golang:alpine
MAINTAINER Dinesh Weerapurage <dweerapurage@mapanything.com>
RUN apk update && apk add git make ca-certificates unzip && update-ca-certificates
# install dep
WORKDIR /root
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src
RUN git clone https://github.com/xydinesh/greenwich.git
WORKDIR greenwich
RUN git checkout dev
RUN make
