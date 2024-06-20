FROM golang:alpine
WORKDIR /build
COPY . .
RUN go build -o /sender cmd/main.go
CMD ["/sender"]
