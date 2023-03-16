# # Set environment variables
# ENV DB_HOST=localhost
# ENV DB_PORT=5432
# ENV DB_USER=myuser
# ENV DB_PASSWORD=mypassword
# ENV DB_NAME=mydb
# ENV PORT=8080
# ENV PGDATA=/var/lib/postgresql/data

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
EXPOSE 3000
CMD ./doxapi