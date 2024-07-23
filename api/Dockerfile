FROM golang:latest

WORKDIR /usr/src/app

COPY . .
RUN rm config.toml
RUN mv config.prod.toml config.toml

RUN go mod download

RUN go build -o main main.go

EXPOSE 2912

CMD ["./main"]
