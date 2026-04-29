FROM golang:1.26.2-alpine3.23

WORKDIR /usr/src/app

COPY Go/go.mod ./
RUN go mod download

COPY Go/. .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
