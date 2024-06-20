FROM golang:alpine
WORKDIR /build
COPY ./cmd/main.go .
RUN go build -o sender main.go
CMD [". /sender"]