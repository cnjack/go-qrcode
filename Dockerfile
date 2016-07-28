FROM golang:1.6.2

MAINTAINER Jack "admin@nightc.com"

RUN go get github.com/cnjack/go-qrcode

WORKDIR $GOPATH/src/github.com/cnjack/go-qrcode

ADD . $GOPATH/src/github.com/cnjack/go-qrcode

RUN godep restore -v

RUN go install -a github.com/cnjack/go-qrcode

RUN cp -r $GOPATH/src/github.com/cnjack/go-qrcode/install/public /go/bin/public
RUN cp -r $GOPATH/src/github.com/cnjack/go-qrcode/install/templates /go/bin/public

EXPOSE 8080

ENTRYPOINT ["/go/bin/go-qrcode"]