FROM golang:1.16
RUN mkdir /app
COPY $PWD /app
WORKDIR /app
RUN go build web-server.go
CMD ["/app/web-server"]
