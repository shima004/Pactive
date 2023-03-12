FROM golang:1.20.1

COPY ./server /go/src/server

WORKDIR /go/src/server

RUN go mod tidy
RUN go mod download

CMD ["go", "run", "main.go"]

