FROM golang:alpine as builder
MAINTAINER Dinesh Weerapurage <dweerapurage@mapanything.com>
RUN apk update && apk add git make ca-certificates unzip curl && update-ca-certificates
# install dep
WORKDIR /root
RUN curl -L -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64
WORKDIR /go/src
RUN chmod 755 /usr/local/bin/dep 
RUN git clone https://github.com/xydinesh/greenwich.git
WORKDIR greenwich
RUN git checkout dev
RUN dep ensure
RUN make

FROM golang:alpine
WORKDIR /root/
COPY --from=builder /go/src/greenwich/bin  /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/greenwich"]