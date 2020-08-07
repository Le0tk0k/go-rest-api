FROM golang:1.14

ENV GO111MODULE=on

WORKDIR /go/src/github.com/Le0tk0k/go-rest-api

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build .

RUN go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]
