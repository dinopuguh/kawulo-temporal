FROM golang:1.14.1

RUN mkdir -p /go/src/github.com/dinopuguh/kawulo-temporal/

WORKDIR /go/src/github.com/dinopuguh/kawulo-temporal/

COPY . .

RUN go build -o temporal main.go

EXPOSE 9090

CMD /go/src/github.com/dinopuguh/kawulo-temporal/temporal