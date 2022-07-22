FROM golang:1.18

WORKDIR /app

COPY . /app

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build cmd/subscriber/main.go

CMD ["/app/main"]