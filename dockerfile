FROM golang:1.20.1

ENV TZ /usr/share/zoneinfo/Asia/Tokyo
ENV ROOT = /go/src/server
ENV GO111MODULE=on

COPY ./server ${ROOT}

WORKDIR ${ROOT}

RUN go mod tidy
RUN go mod download

RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]

