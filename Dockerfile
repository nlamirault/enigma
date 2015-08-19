FROM golang:1.3-cross
RUN apt-get update && apt-get install -y --no-install-recommends openssh-client
RUN go get github.com/mitchellh/gox
RUN go get github.com/aktau/github-release

ENV GOPATH /go/
WORKDIR /go/src/github.com/nlamirault/enigma

ADD . /go/src/github.com/nlamirault/enigma
