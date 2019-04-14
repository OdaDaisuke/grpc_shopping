FROM golang:1.9

RUN go get -u gopkg.in/godo.v2/cmd/godo

WORKDIR /go/src/app

EXPOSE 8080

CMD ["/go/bin/godo", "server", "--watch"]