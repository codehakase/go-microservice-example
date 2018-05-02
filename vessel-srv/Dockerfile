FROM golang:1.10.0 as builder

WORKDIR /home/codehakase/gocode/src/github.com/codehakase/go-microservice-example/vessel-srv

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM debian:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /home/codehakase/gocode/src/github.com/codehakase/go-microservice-example/vessel-srv/vessel-srv .

CMD ["./vessel-srv"]
