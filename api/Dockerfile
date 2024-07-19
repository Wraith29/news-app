FROM golang:latest

WORKDIR /usr/src/app

COPY go.mod go.sum .

RUN go mod download

COPY *.go .

RUN go build -o main main.go

EXPOSE 2912

CMD ["./main"]
