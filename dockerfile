FROM golang:1.14-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o http-grpc-api ./cmd/main.go

CMD ["./http-grpc-api"]