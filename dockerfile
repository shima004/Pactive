FROM golang:1.20.1

COPY ./server /go/src/server

WORKDIR /go/src/server

RUN go mod tidy

CMD ["go", "run", "main.go"]

