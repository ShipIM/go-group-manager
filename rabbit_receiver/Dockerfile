FROM golang:alpine
WORKDIR /build
COPY . .
RUN go build -o /rabbit-receiver cmd/main.go
CMD ["/rabbit-receiver"]
