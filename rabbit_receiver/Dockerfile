FROM golang:alpine
WORKDIR /build
COPY ./cmd/main.go .
RUN go build -o rabbit-receiver main.go
CMD [". /rabbit-receiver"]