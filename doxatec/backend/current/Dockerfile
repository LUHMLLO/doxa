#build stage
FROM golang:latest as builder
ENV GOOS linux
ENV CGO_ENABLED 0
WORKDIR /app
COPY . .
RUN go build -o ./doxapi

#production stage
FROM alpine:latest as production
RUN apk add --no-cache ca-certificates
COPY --from=builder app .
EXPOSE 8081
CMD ./doxapi