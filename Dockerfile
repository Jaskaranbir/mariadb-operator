FROM golang:1.13.0-stretch
LABEL maintainer="Jaskaranbir Dhillon"

RUN mkdir /app
WORKDIR /app

ADD . .
RUN go build -o main .

CMD ["/app/main"]
