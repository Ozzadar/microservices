FROM golang:1.11.0 as builder

WORKDIR /go/src/github.com/ozzadar/microservices/email-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o email-service -a -installsuffix cgo main.go

FROM debian:latest

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/ozzadar/microservices/email-service .

CMD ["./email-service"]