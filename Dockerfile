FROM golang:1.13.7

WORKDIR /go/src/rabbitmkube
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["rabbitmkube"]