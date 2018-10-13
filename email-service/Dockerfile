
FROM golang:1.11.0 as builder

WORKDIR /go/src/github.com/ozzadar/microservices/email-service

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM debian:latest

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/ozzadar/microservices/email-service .

CMD ["./email-service"]