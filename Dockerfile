FROM golang:latest

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go build -o main cmd/api/main.go
CMD ["/app/main"]