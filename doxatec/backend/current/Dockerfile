#build stage
FROM golang:1.20.1 as builder

ENV GOOS linux
ENV CGO_ENABLED 0

WORKDIR /app

COPY . .

RUN go build -o ./doxapi

EXPOSE 8081

CMD ./doxapi