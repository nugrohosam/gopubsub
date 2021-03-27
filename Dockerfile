FROM golang:1.6

ARG TYPE=subscriber
ARG LISTEN_PORT=1234
ARG LISTEN_HOST=localhost

WORKDIR /go/src/app
COPY ./$TYPE .
COPY go.mod .
COPY go.sum .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["main", "--host", $LISTEN_HOST, "--listen", $LISTEN_PORT]