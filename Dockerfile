FROM golang:1.20

WORKDIR /go/goapp

RUN apt-get update && apt-get install -y librdkafka-dev sqlite3

CMD ["tail", "-f", "/dev/null"]