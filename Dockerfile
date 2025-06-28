FROM docker.arvancloud.ir/golang:1.24 as builder

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download


COPY ./ ./

RUN go build -o app ./cmd/mainservice/main.go

EXPOSE 5000

CMD ["./app"]