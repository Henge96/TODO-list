FROM golang:1.18.2-alpine3.15

ENV GOPATH/=
WORKDIR /app

COPY . /app

RUN go build ./cmd/main.go

EXPOSE 8000

CMD ./main
