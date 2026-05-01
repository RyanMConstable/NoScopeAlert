FROM golang:1.26.2-alpine3.23

WORKDIR /usr/src/downloadmatches

COPY Go/go.mod Go/go.sum ./
RUN go mod download

COPY Go/. .
RUN go build -v -o /usr/local/bin/downloadmatches ./...

CMD ["downloadmatches"]
