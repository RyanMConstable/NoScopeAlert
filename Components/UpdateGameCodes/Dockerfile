FROM golang:1.26.2-alpine3.23

WORKDIR /usr/src/updategamecodes

COPY Go/go.mod Go/go.sum ./
RUN go mod download

COPY Go/. .
RUN go build -v -o /usr/local/bin/updategamecodes ./...

CMD ["updategamecodes"]
