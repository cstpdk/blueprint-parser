FROM golang

ADD . /go/src/github.com/cstpdk/blueprint-parser

WORKDIR /go/src/github.com/cstpdk/blueprint-parser

RUN go get -t ./...
RUN go install

ENTRYPOINT ["blueprint-parser"]
