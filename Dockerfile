FROM golang:1.12.9

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]

